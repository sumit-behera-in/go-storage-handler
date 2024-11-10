package db

import (
	"encoding/json"
	"os"
)

func (c *Clients) UpdateJson(jsonPath string) {
	println("update json called on ")
	i := 0
	for i < len(c.dbCollection.Database) {
		c.dbCollection.Database[i].UsedSpaceGB = c.clients[i].updateSpace()
		i++
	}

	updatedData, err := json.MarshalIndent(c.dbCollection, "", "  ")
	if err != nil {
		println("Error marshaling JSON:", err)
		return
	}

	if err := os.WriteFile(jsonPath, updatedData, 0644); err != nil {
		println("Error writing to file:", err)
		return
	}
	println("JSON file updated successfully.")
}
