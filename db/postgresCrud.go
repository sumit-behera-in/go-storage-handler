package db

import (
	"fmt"
	"log"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/util"
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

	return pc.db.QueryRow(query, data.fileName, data.File).Err()
}

func (pc *postgresClient) download(fileName string, fileType string) {
	data := Data{}
	data.fileName = fileName
	query := fmt.Sprintf("SELECT file FROM %s WHERE file_name = $1", fileType)
	pc.db.QueryRow(query, fileName).Scan(&data.File)

	if !data.isEmpty() {
		downloadPath, err := util.GetDefaultDownloadPath()
		if err != nil {
			log.Fatal("Error writing file:", err)
		}

		outputPath := fmt.Sprintf("%s/%s", downloadPath, fileName)
		err = os.WriteFile(outputPath, data.File, 0666)
		if err != nil {
			log.Fatal("Error writing file:", err)
		}
		fmt.Printf("File %s downloaded successfully to %s\n", fileName, downloadPath)
	}
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

func (pc *postgresClient) updateSpace() float64 {
	var totalSizeBytes int64
	err := pc.db.Get(&totalSizeBytes, `SELECT pg_database_size(current_database())`)
	if err != nil {
		log.Fatal(err)
	}

	// Convert bytes to GB
	totalSizeGB := float64(totalSizeBytes) / (1024 * 1024 * 1024)
	return totalSizeGB
}

func (pc *postgresClient) find(fileName string, fileType string) bool {
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %v WHERE file_name = $1);`, fileType)
	err := pc.db.Get(&exists, query, fileName)
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}

	return exists
}
