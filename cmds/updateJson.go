package cmds

import (
	"encoding/json"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/db"
)

func UpdateJson(c db.Clients, jsonPath string) {
	println("update json called on ")
	i := 0
	for i < len(c.Clients) {
		c.DBCollection.Database[i].UsedSpaceGB = c.Clients[i].UpdateSpace()
		i++
	}

	updatedData, err := json.MarshalIndent(c.DBCollection, "", "  ")
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
