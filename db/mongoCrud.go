package db

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mc *mongoClient) insertRow(data Data) error {
	bucket, err := gridfs.NewBucket(mc.database, options.GridFSBucket().SetName(data.FileType))
	if err != nil {
		return err
	}

	// uploadstream is used to upload the data
	uploadStream, err := bucket.OpenUploadStreamWithID(nil, data.FileName)
	if err != nil {
		return fmt.Errorf("failed to open upload stream: %v", err)
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(data.File)
	if err != nil {
		return fmt.Errorf("failed to upload file data: %v", err)
	}

	fmt.Printf("File %s uploaded successfully\n", data.FileName)
	return nil
}

func (mc *mongoClient) download(filename string, tableName string) {
	data := Data{}

	// Create a GridFS bucket
	bucket, err := gridfs.NewBucket(mc.database)
	if err != nil {
		log.Fatal("failed to create GridFS bucket:", err)
	}

	filter := bson.M{"name": filename}
	err = mc.database.Collection(tableName+".files").FindOne(mc.ctx, filter).Decode(&data)
	if err != nil {
		log.Fatal("Error retrieving file from database:", err)
	}

	// Open a download stream
	downloadStream, err := bucket.OpenDownloadStreamByName(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer downloadStream.Close()

	if !data.isEmpty() {
		// Create a new file to save the downloaded data
		outFile, err := os.Create(util.DOWNLOAD_PATH + "/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()

		// Read from GridFS and write to file
		if _, err := io.Copy(outFile, downloadStream); err != nil {
			log.Fatalf("failed to write data to file: %v", err)
		}

		fmt.Printf("File %s downloaded successfully to %s\n", filename, util.DOWNLOAD_PATH)

	}
}
