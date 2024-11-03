package db

import (
	"fmt"
)

func (pc *postgresClient) createTable(tableName string) error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		file BYTEA NOT NULL,		
		created timestamp DEFAULT NOW()
	)`, tableName)

	_, err := pc.db.Exec(query)
	return err
}

func (pc *postgresClient) insertRowQuery(data Data) error {
	query := fmt.Sprintf(`INSERT INTO %s (file_name,file)
	VALUES ($1,$2) RETURNING id`, data.FileType)

	return pc.db.QueryRow(query, data.FileName, data.File).Err()
}
