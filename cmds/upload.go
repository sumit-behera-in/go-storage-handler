package cmds

import "github.com/urfave/cli/v2"

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
		Action: upload,
	}
}

func upload(*cli.Context) error {

	return nil

}
