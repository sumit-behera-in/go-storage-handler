package cmds

import (
	"github.com/urfave/cli/v2"
)

func Download() *cli.Command {
	return &cli.Command{
		Name:  "download",
		Usage: "downloads a file to the storage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Usage:    "downlod a file to our database",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			data := ctx.String("file")
			Clients.Download(data)
			return nil
		},
	}
}
