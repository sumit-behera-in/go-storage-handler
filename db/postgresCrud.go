package db

import (
	"fmt"
	"log"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/util"
)

func (pc *postgresClient) createTable(tableName string) error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		file BYTEA NOT NULL
	)`, tableName)

	_, err := pc.db.Exec(query)
	return err
}

func (pc *postgresClient) insertRowQuery(data Data) error {
	query := fmt.Sprintf(`INSERT INTO %s (name,file)
	VALUES ($1,$2) RETURNING id`, data.FileType)

	return pc.db.QueryRow(query, data.FileName, data.File).Err()
}

func (pc *postgresClient) download(filename string, tableName string) {
	data := Data{}
	query := fmt.Sprintf("SELECT name, file FROM %s WHERE id = $1", tableName)
	err := pc.db.QueryRow(query, filename).Scan(&data.FileName, &data.File)

	if err != nil {
		log.Fatal("Error retrieving file from database:", err)
	}

	if !data.isEmpty() {
		outputPath := fmt.Sprintf("%s/%s", util.DOWNLOAD_PATH, filename)
		err = os.WriteFile(outputPath, data.File, 0666)
		if err != nil {
			log.Fatal("Error writing file:", err)
		}
	}
}
