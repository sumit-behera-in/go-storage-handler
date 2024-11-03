package main

import (
	"log"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/cmds"
	"github.com/sumit-behera-in/go-storage-handler/db"
	"github.com/urfave/cli/v2"
)

var clients db.Clients

func main() {
	app := &cli.App{
		Name:  "Go Storage Handler",
		Usage: "It is used to handle multiple storages",
		Commands: []*cli.Command{
			cmds.Upload(clients),
			cmds.Update(clients),
			cmds.Download(clients),
			cmds.Delete(clients),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "configPath",
				Usage:    "contains config.json file",
				Value:    "config.json",
				Required: true,
			},
		},
		Before: func(ctx *cli.Context) error {
			filePath := ctx.String("configPath")
			var err error
			clients, err = db.New(filePath)
			return err
		},
		Action: func(ctx *cli.Context) error {
			return nil
		},
		After: func(ctx *cli.Context) error {
			clients.Close()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
