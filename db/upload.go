package db

import (
	"fmt"
)

func (c *Clients) Upload(data Data, sizeOfTheData float64) error {

	// check if the data exists anywhere
	var exists bool
	for i, client := range c.Clients {
		exists = client.find(data.FileName, data.FileType)
		if exists {
			return fmt.Errorf("the file already exists on database with index : %v", i)
		}
	}

	var client client
	var space_available = false

	i := 0

	for i < len(c.DBCollection.Database) {
		c.DBCollection.Database[i].UsedSpaceGB = c.Clients[i].UpdateSpace()
		if space_available = c.isAvailspace(sizeOfTheData, i); space_available {
			break
		}
		i++
	}

	if !space_available {
		return fmt.Errorf("unable to upload as any of the databases cant hold this data")
	}

	client = c.Clients[i]

	err := client.upload(data)
	if err != nil {
		return err
	}

	println("Successfully uploaded to database index: %v", i)
	return nil

}

func (c *Clients) isAvailspace(data float64, index int) bool {
	return (c.DBCollection.Database[index].UsedSpaceGB + data) <= 0.8*c.DBCollection.Database[index].TotalSpaceGB
}
