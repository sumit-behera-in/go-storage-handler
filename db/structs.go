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
	Protocol      string `json:"protocol,omitempty"`
	ConnectionURL string `json:"connetionURL"`
	Port          int    `json:"port"`
	User          string `json:"user,omitempty"`
	Password      string `json:"password,omitempty"`
	DBName        string `json:"dbName"`
}

type Database struct {
	Priority     int     `json:"priority"`
	TotalSpaceGB float64 `json:"total_space_GB"`
	UsedSpaceGB  float64 `json:"used_space_GB"`
	DBProvider   string  `json:"db_provider"`
	Config       Config  `json:"config"`
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

func (d *Data) isEmpty() bool {
	return len(d.File) == 0
}
