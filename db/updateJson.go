package db

import (
	"encoding/json"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/util"
	"github.com/urfave/cli/v2"
)

func (clients *Clients) UpdateJson(ctx *cli.Context) {
	jsonPath := ctx.String(util.ConfigPath)
	updatedData, err := json.MarshalIndent(clients, "", "  ")
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
