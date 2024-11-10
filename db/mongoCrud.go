package db

import (
	"fmt"
	"io"
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
	data := Data{
		fileName: fileName,
		FileType: fileType,
	}

	bucket, err := gridfs.NewBucket(mc.database, options.GridFSBucket().SetName(fileType))
	if err != nil {
		log.Fatalf("failed to create GridFS bucket: %v", err)
	}

	// Find the file with the given fileName.
	var fileID primitive.ObjectID
	cursor, err := bucket.Find(bson.M{"filename": fileName})
	if err != nil {
		log.Fatalf("failed to find file: %v", err)
	}
	defer cursor.Close(mc.ctx)

	if cursor.Next(mc.ctx) {
		var fileInfo bson.M
		if err = cursor.Decode(&fileInfo); err != nil {
			log.Fatalf("failed to decode file info: %v", err)
		}
		fileID = fileInfo["_id"].(primitive.ObjectID)
	} else {
		log.Fatalf("file %s not found in bucket %s", fileName, fileType)
	}

	// Open the download stream with the located fileID.
	downloadStream, err := bucket.OpenDownloadStream(fileID)
	if err != nil {
		log.Fatalf("failed to open download stream: %v", err)
	}
	defer downloadStream.Close()

	// Read the file data from the download stream into a byte slice.
	fileData := make([]byte, downloadStream.GetFile().Length)
	_, err = downloadStream.Read(fileData)
	if err != nil && err != io.EOF {
		log.Fatalf("failed to read file data: %v", err)
	}

	// Store the file data in the data struct.
	data.File = fileData

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
	bucket, err := gridfs.NewBucket(mc.database, options.GridFSBucket().SetName(fileType))
	if err != nil {
		return fmt.Errorf("failed to create GridFS bucket: %v", err)
	}

	// Find the file with the given fileName.
	var fileID primitive.ObjectID
	cursor, err := bucket.Find(bson.M{"filename": fileName})
	if err != nil {
		return fmt.Errorf("failed to find file: %v", err)
	}
	defer cursor.Close(mc.ctx)

	// Retrieve the file's ObjectID from the cursor.
	if cursor.Next(mc.ctx) {
		var fileInfo bson.M
		if err = cursor.Decode(&fileInfo); err != nil {
			return fmt.Errorf("failed to decode file info: %v", err)
		}
		fileID = fileInfo["_id"].(primitive.ObjectID)
	} else {
		return fmt.Errorf("file %s not found in bucket %s", fileName, fileType)
	}

	// Delete the file using its ObjectID.
	if err = bucket.Delete(fileID); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	fmt.Printf("File %s successfully deleted from bucket %s\n", fileName, fileType)
	return nil
}

func (mc *mongoClient) updateSpace() float64 {
	var result bson.M
	if err := mc.database.RunCommand(mc.ctx, bson.D{{Key: "dbStats", Value: 1}}).Decode(&result); err != nil {
		log.Fatalf("failed to run dbStats command: %v", err)
	}

	// Retrieve the storage size in bytes, handling both int64 and float64 cases
	var storageSizeBytes float64
	switch size := result["storageSize"].(type) {
	case int32:
		storageSizeBytes = float64(size)
	case int64:
		storageSizeBytes = float64(size)
	case float64:
		storageSizeBytes = size
	default:
		log.Fatalf("failed to retrieve storageSize: unexpected type %T", size)
	}

	// Convert bytes to gigabytes
	sizeInGB := storageSizeBytes / (1024 * 1024 * 1024)
	return sizeInGB
}
