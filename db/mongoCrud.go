package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (mc *mongoClient) insertRow(data Data) error {
	bucket, err := gridfs.NewBucket(mc.database)
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
