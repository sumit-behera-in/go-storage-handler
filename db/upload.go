package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sumit-behera-in/go-storage-handler/util"
)

func (c *Clients) Upload(fPath string) error {

	var client client
	var space_available = false

	sizeOfTheData, err := getFileSizeGB(fPath)
	if err != nil {
		return nil
	}

	i := 0

	for i < len(c.dbCollection.Database) {
		c.dbCollection.Database[i].UsedSpaceGB = c.clients[i].updateSpace()
		if space_available = c.isAvailspace(sizeOfTheData, i); space_available {
			break
		}
		i++
	}

	if !space_available {
		return fmt.Errorf("unable to upload as any of the databases cant hold this data")
	}

	client = c.clients[i]

	var data = Data{}

	data.fileName = filepath.Base(fPath)

	// Get the file extension
	fileExtension := filepath.Ext(fPath)

	// Remove the leading dot from the extension, if it exists
	if len(fileExtension) > 0 {
		data.FileType = fileExtension[1:]
	} else {
		data.FileType = util.UNKNOWN_FILE_TYPE
	}

	data.File, err = os.ReadFile(fPath)
	if err != nil {
		return fmt.Errorf("error while creating a table %v", err)
	}

	err = client.upload(data)
	if err != nil {
		return err
	}

	println("Sucessfully uploaded to database index: %v", i)
	return nil

}

func getFileSizeGB(filePath string) (float64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	// Get size in bytes and convert to gigabytes
	fileSizeGB := float64(fileInfo.Size()) / (1024 * 1024 * 1024)
	return fileSizeGB, nil
}

func (c *Clients) isAvailspace(data float64, index int) bool {
	return (c.dbCollection.Database[index].UsedSpaceGB + data) <= 0.8*c.dbCollection.Database[index].TotalSpaceGB
}
