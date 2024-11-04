package main

import (
	"log"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/cmds"
	"github.com/sumit-behera-in/go-storage-handler/db"
	"github.com/sumit-behera-in/go-storage-handler/util"
	"github.com/urfave/cli/v2"
)


func main() {
	app := &cli.App{
		Name:  "Go Storage Handler",
		Usage: "It is used to handle multiple storages",
		Commands: []*cli.Command{
			cmds.Upload(),
			// cmds.Update(clients),
			// cmds.Download(clients),
			// cmds.Delete(clients),
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
			cmds.Clients, err = db.New(filePath)
			return err
		},
		Action: func(ctx *cli.Context) error {
			return nil
		},
		After: func(ctx *cli.Context) error {
			cmds.Clients.Close()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
