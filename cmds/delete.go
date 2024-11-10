package cmds

import (
	"github.com/urfave/cli/v2"
)

func Delete() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "delete a file to the storage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Usage:    "name of the file to delete from our database",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			filePath := ctx.String("file")
			Clients.Delete(filePath)
			return nil
		},
	}
}
