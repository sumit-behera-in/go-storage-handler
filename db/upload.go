package db

import (
	"fmt"
	"os"
)

func (c *Clients) Upload(data string) error {

	var client client
	var space_available = false

	sizeOfTheData, err := getFileSizeGB(data)
	if err != nil {
		return nil
	}

	i := 0

	for _, clientData := range c.clients {
		client = clientData
		if space_available = client.availspace(sizeOfTheData); space_available == true {
			break
		}
		i++
	}

	if !space_available {
		return fmt.Errorf("unable to upload as any of the databases cant hold this data")
	}

	err = client.upload(data)
	if err != nil {
		return err
	}

	println("Sucessfully uploaded to database No: %v", i)
	client.updateSpace(sizeOfTheData)
	return nil

}

func (mc *mongoClient) upload(data string) error {
	return nil
}

func (pc *postgresClient) upload(data string) error {
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

func (mc *mongoClient) availspace(data float64) bool {
	return (mc.availableSpaceGB + data) <= 0.8*mc.totalSpaceGB
}

func (pc *postgresClient) availspace(data float64) bool {
	return (data + pc.availableSpaceGB) <= 0.8*pc.totalSpaceGB
}

func (mc *mongoClient) updateSpace(data float64) {
	mc.availableSpaceGB += data
}

func (pc *postgresClient) updateSpace(data float64) {
	pc.availableSpaceGB += data
}
