package db

import (
	"fmt"
	"log"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mc *mongoClient) upload(data Data) error {
	// Create a new GridFS bucket with the specified file type.
	bucket, err := gridfs.NewBucket(mc.database, options.GridFSBucket().SetName(data.FileType))
	if err != nil {
		return err
	}

	// Generate a new ObjectID to use as the file's ID.
	fileID := primitive.NewObjectID()

	// Open an upload stream with the generated file ID.
	uploadStream, err := bucket.OpenUploadStreamWithID(fileID, data.fileName)
	if err != nil {
		return fmt.Errorf("failed to open upload stream: %v", err)
	}
	defer uploadStream.Close()

	// Write the file data to the upload stream.
	_, err = uploadStream.Write(data.File)
	if err != nil {
		return fmt.Errorf("failed to upload file data: %v", err)
	}

	fmt.Printf("File %s uploaded successfully with ID %s\n", data.fileName, fileID.Hex())
	return nil
}

func (mc *mongoClient) download(fileName string, fileType string) {
	data := Data{}

	// Create a GridFS bucket
	println("dtatabase :", mc.database)
	bucket, err := gridfs.NewBucket(mc.database, options.GridFSBucket().SetName(fileType))
	if err != nil {
		log.Fatal("failed to create GridFS bucket:", err)
	}

	filter := bson.M{"name": fileName}
	println(fileType + ".files")
	err = mc.database.Collection(fileType+".files").FindOne(mc.ctx, filter).Decode(&data)
	if err != nil {
		log.Fatal("Error retrieving file from database:", err)
	}

	// Open a download stream
	downloadStream, err := bucket.OpenDownloadStreamByName(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer downloadStream.Close()

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

func (mc *mongoClient) delete(fileName string, fileType string) error {
	return nil
}
