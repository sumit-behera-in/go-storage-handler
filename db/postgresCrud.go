package db

import (
	"fmt"
	"log"
)

func (pc *postgresClient) createTable(fileType string) error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		file_name VARCHAR(100) NOT NULL UNIQUE,
		file BYTEA NOT NULL
	)`, fileType)

	_, err := pc.db.Exec(query)
	return err
}

func (pc *postgresClient) upload(data Data) error {
	err := pc.createTable(data.FileType)
	if err != nil {
		return fmt.Errorf("error while creating a table %v", err)
	}
	query := fmt.Sprintf(`INSERT INTO %s (file_name,file)
	VALUES ($1,$2) RETURNING id`, data.FileType)

	return pc.db.QueryRow(query, data.FileName, data.File).Err()
}

func (pc *postgresClient) download(fileName string, fileType string) Data {
	data := Data{}
	data.FileName = fileName
	data.FileType = fileType
	query := fmt.Sprintf("SELECT file FROM %s WHERE file_name = $1", fileType)
	err := pc.db.QueryRow(query, fileName).Scan(&data.File)
	if err != nil {
		log.Printf("Download Failed for %s with error : %w", fileName, err)		
	}

	return data
}

func (pc *postgresClient) delete(fileName string, fileType string) error {

	query := fmt.Sprintf(
		`DELETE FROM %s
        WHERE file_name = $1;`,
		fileType,
	)

	_, err := pc.db.Exec(query, fileName)
	if err != nil {
		return err
	}

	vaccumSQL := fmt.Sprintf("VACCUM FULL %s", fileType)
	_, err = pc.db.Exec(vaccumSQL)

	return err
}

func (pc *postgresClient) UpdateSpace() float64 {
	var totalSizeBytes int64
	err := pc.db.Get(&totalSizeBytes, `SELECT pg_database_size(current_database())`)
	if err != nil {
		log.Printf("Error while updating space in postgres error : %w", err)
	}

	// Convert bytes to GB
	totalSizeGB := float64(totalSizeBytes) / (1024 * 1024 * 1024)
	return totalSizeGB
}

func (pc *postgresClient) find(fileName string, fileType string) bool {
	pc.createTable(fileType)

	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %v WHERE file_name = $1);`, fileType)
	err := pc.db.Get(&exists, query, fileName)
	if err != nil {
		log.Printf("failed to execute query: %w", err)
	}

	return exists
}
