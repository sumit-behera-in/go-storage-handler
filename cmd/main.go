package main

import (
	"log"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/db"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "Go Storage Handler",
		Usage:    "It is used to handle multiple storages",
		Commands: []*cli.Command{},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "configPath",
				Usage:    "contains config.json file",
				Value:    "config.json",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := db.New(ctx)
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
