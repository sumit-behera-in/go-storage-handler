package cmds

import (
	"github.com/sumit-behera-in/go-storage-handler/util"
	"github.com/urfave/cli/v2"
)

func Upload() *cli.Command {
	return &cli.Command{
		Name:  "upload",
		Usage: "uploads a file to the storage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Usage:    "uploads a file to our database",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			data := ctx.String("file")
			jsonPath := ctx.String(util.ConfigPath)
			defer Clients.UpdateJson(jsonPath)
			return Clients.Upload(data)
		},
	}
}
