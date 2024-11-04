package db

import (
	"encoding/json"
	"os"
)

func (c *Clients) UpdateJson(jsonPath string) {
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
