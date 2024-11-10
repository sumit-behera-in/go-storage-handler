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
	err := pc.db.QueryRow(query, fileName).Scan(&data.File)

	if err != nil {
		log.Fatal("Error retrieving file from database:", err)
	}

	if !data.isEmpty() {
		outputPath := fmt.Sprintf("%s/%s", util.DOWNLOAD_PATH, fileName)
		err = os.WriteFile(outputPath, data.File, 0666)
		if err != nil {
			log.Fatal("Error writing file:", err)
		}
		fmt.Printf("File %s downloaded successfully to %s\n", fileName, util.DOWNLOAD_PATH)
	}
}

// func (pc *postgresClient) update(data Data) error {
// 	query := fmt.Sprintf(
// 		`INSERT INTO %v (file_name,file)
// 		VALUES ($1,$2)
// 		ON CONFLICT (file_name)
// 		DO UPDATE SET
// 		file = EXCLUDED.file
// 		RETURNING id`, data.FileType,
// 	)
// 	_, err := pc.db.Exec(query, data.fileName, data.FileType)
// 	return err
// }

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

	vaccumSQL := fmt.Sprintf("VACCUM %s", fileType)
	_, err = pc.db.Exec(vaccumSQL)

	return err
}
