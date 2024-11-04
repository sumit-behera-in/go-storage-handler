package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoClient struct {
	db      *mongo.Client
	ctx     context.Context
	closeDB context.CancelFunc
}

type postgresClient struct {
	db *sqlx.DB
}

type client interface {
	connect(Database) error
	upload(string) error
	close()
}

type Clients struct {
	clients      []client
	dbCollection DBCollection
}

type config struct {
	Protocol      string `json:"protocol,omitempty"`
	ConnectionURL string `json:"connetionURL"`
	Port          int    `json:"port"`
	User          string `json:"user,omitempty"`
	Password      string `json:"password,omitempty"`
	DBName        string `json:"dbName"`
}

type Database struct {
	Priority         int     `json:"priority"`
	TotalSpaceGB     float64 `json:"total_space_GB"`
	AvailableSpaceGB float64 `json:"available_space_GB"`
	DBProvider       string  `json:"db_provider"`
	Config           config  `json:"config"`
}

type DBCollection struct {
	Project  string     `json:"project"`
	Database []Database `json:"database"`
}

type Data struct {
	FileName string
	FileType string
	File     []byte
}
