package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli/v2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/lib/pq"
)

func New(ctx *cli.Context) (Clients, error) {

	clients := Clients{}
	filePath := ctx.String("configPath")

	file, err := os.ReadFile(filePath)
	if err != nil {
		return clients, err
	}

	// Unmarshal the JSON into the DBBCollection struct
	var dbBCollection DBCollection
	if err := json.Unmarshal(file, &dbBCollection); err != nil {
		return clients, fmt.Errorf("error parsing json: %v", err)
	}

	for i, db := range dbBCollection.Database {
		if err = clients.addConnect(db.DBProvider, db); err != nil {
			return clients, fmt.Errorf("error connecting client no.: %v , err %v", i, err)
		}
	}

	return clients, nil
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
		c.clients = append(c.clients, mongoClient)
	} else if dbProvider == "postgres" {
		err = postgresClient.connect(db)
		if err != nil {
			return err
		}
		c.clients = append(c.clients, postgresClient)

	}
	return err
}

func (client *mongoClient) connect(db Database) error {

	var err error

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s%s", db.Config.Protocol, db.Config.ConnectionURL))
	client.ctx, client.close = context.WithTimeout(context.Background(), 5*time.Minute)
	client.db, err = mongo.Connect(client.ctx, clientOptions)

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
		"password=%s sslmode=disable",
		db.Config.ConnectionURL, db.Config.Port, db.Config.User, db.Config.Password)

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
