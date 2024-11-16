package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/lib/pq"
)

func New(dbBCollection DBCollection) (Clients, error) {
	Clients := Clients{}

	Clients.DBCollection = dbBCollection

	for i, db := range dbBCollection.Database {
		if err := Clients.addConnect(db.DBProvider, db); err != nil {
			return Clients, fmt.Errorf("error connecting client no.: %v , err %w", i, err)
		}
	}

	return Clients, nil
}

func (c *Clients) addConnect(dbProvider string, db Database) error {

	mongoClient := &mongoClient{}
	postgresClient := &postgresClient{}
	var err error
	if dbProvider == "mongodb" {
		err = mongoClient.connect(db)
		if err != nil {
			return err
		}
		c.Clients = append(c.Clients, mongoClient)
	} else if dbProvider == "postgres" {
		err = postgresClient.connect(db)
		if err != nil {
			return err
		}
		c.Clients = append(c.Clients, postgresClient)

	}
	return err
}

func (client *mongoClient) connect(db Database) error {

	var err error

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s%s", db.Config.Protocol, db.Config.ConnectionURL))
	client.ctx, client.closeDB = context.WithTimeout(context.Background(), 5*time.Minute)
	client.db, err = mongo.Connect(client.ctx, clientOptions)
	client.database = client.db.Database(db.Config.DBName)

	if err != nil {
		return err
	}

	err = client.db.Ping(client.ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	return err
}

func (client *postgresClient) connect(db Database) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Config.ConnectionURL, db.Config.Port, db.Config.User, db.Config.Password, db.Config.DBName)

	var err error
	client.db, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = client.db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected To postgres")

	return err
}

func (client *postgresClient) close() {
	client.db.Close()
}

func (client *mongoClient) close() {
	client.closeDB()
}

func (c *Clients) Close() {
	for _, client := range c.Clients {
		client.close()
	}
}
