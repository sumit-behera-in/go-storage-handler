package cmds

import (
	"github.com/urfave/cli/v2"
)

func Update() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "update a file to the storage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Usage:    "file path to update a file to our database",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			data, sizeOfTheData := getInfo(ctx)
			return Clients.Update(data, sizeOfTheData)
		},
	}
}
