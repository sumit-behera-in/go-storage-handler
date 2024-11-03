package cmds

import (
	"fmt"

	"github.com/sumit-behera-in/go-storage-handler/db"
	"github.com/urfave/cli/v2"
)

func Upload(clients db.Clients) *cli.Command {
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
			fmt.Println("upload")
			fmt.Println(data)
			return clients.Upload(data)
		},
	}
}
