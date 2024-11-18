package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoClient struct {
	db       *mongo.Client
	ctx      context.Context
	closeDB  context.CancelFunc
	database *mongo.Database
}

type postgresClient struct {
	db *sqlx.DB
}

type Client interface {
	connect(Database) error
	upload(Data) error
	download(string, string)
	delete(string, string) error
	UpdateSpace() float64
	find(string, string) bool
	close()
}

type Clients struct {
	Clients      []Client
	DBCollection DBCollection
}

type Config struct {
	Protocol      string `bson:"protocol,omitempty" json:"protocol,omitempty"`
	ConnectionURL string `bson:"connectionURL" json:"connectionURL"`
	Port          int    `bson:"port" json:"port"`
	User          string `bson:"user,omitempty" json:"user,omitempty"`
	Password      string `bson:"password,omitempty" json:"password,omitempty"`
	DBName        string `bson:"dbName" json:"dbName"`
}

type Database struct {
	DBProvider   string  `bson:"db_provider" json:"db_provider"`
	Priority     int     `bson:"priority" json:"priority"`
	TotalSpaceGB float64 `bson:"total_space_GB" json:"total_space_GB"`
	UsedSpaceGB  float64 `bson:"used_space_GB" json:"used_space_GB"`
	Config       Config  `bson:"config" json:"config"`
}

type DBCollection struct {
	Project  string     `bson:"project" json:"config"`
	Database []Database `bson:"database" json:"database"`
}

type Data struct {
	FileName string
	FileType string
	File     []byte
}

func (d *Data) isEmpty() bool {
	return len(d.File) == 0
}
