package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/cmds"
	"github.com/sumit-behera-in/go-storage-handler/db"
	"github.com/sumit-behera-in/go-storage-handler/util"
	"github.com/urfave/cli/v2"
)

const (
	version = "1.0.0"
)

func main() {
	app := &cli.App{
		Name:    "Go Storage Handler",
		Usage:   "It is used to handle multiple storages",
		Version: version,
		Commands: []*cli.Command{
			cmds.Upload(),
			cmds.Update(),
			cmds.Download(),
			cmds.Delete(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     util.ConfigPath,
				Usage:    "contains config.json file",
				Value:    "config.json",
				Required: true,
			},
		},
		Before: func(ctx *cli.Context) error {
			filePath := ctx.String(util.ConfigPath)
			var err error

			file, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}

			// Unmarshal the JSON into the DBBCollection struct
			var dbBCollection db.DBCollection
			if err := json.Unmarshal(file, &dbBCollection); err != nil {
				return fmt.Errorf("error parsing json: %v", err)
			}
			cmds.Clients, err = db.New(dbBCollection)
			return err
		},
		Action: func(ctx *cli.Context) error {
			return nil
		},
		After: func(ctx *cli.Context) error {
			jsonPath := ctx.String(util.ConfigPath)
			cmds.UpdateJson(cmds.Clients, jsonPath)
			cmds.Clients.Close()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
