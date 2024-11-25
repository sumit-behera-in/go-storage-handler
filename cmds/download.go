package cmds

import (
	"fmt"
	"os"

	"github.com/sumit-behera-in/go-storage-handler/util"
	"github.com/urfave/cli/v2"
)

func Download() *cli.Command {
	return &cli.Command{
		Name:  "download",
		Usage: "downloads a file to the storage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Usage:    "provide the name of th file to be downloded from our database",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			fileName := ctx.String("file")
			data := Clients.Download(fileName)

			if !data.IsEmpty() {
				downloadPath, err := util.GetDefaultDownloadPath()
				if err != nil {
					return err
				}

				outputPath := fmt.Sprintf("%s/%s", downloadPath, fileName)
				err = os.WriteFile(outputPath, data.File, 0666)
				if err != nil {
					return err
				}
				fmt.Printf("File %s downloaded successfully to %s\n", fileName, downloadPath)
			}
			return nil
		},
	}
}
