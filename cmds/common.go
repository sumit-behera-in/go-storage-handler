package cmds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sumit-behera-in/go-storage-handler/db"
	"github.com/sumit-behera-in/go-storage-handler/util"
	"github.com/urfave/cli/v2"
)

var (
	Clients db.Clients
)

func getInfo(ctx *cli.Context) (db.Data, float64) {
	fPath := ctx.String("file")
	var data = db.Data{}
	sizeOfTheData, err := getFileSizeGB(fPath)
	if err != nil {
		panic(err)
	}
	data.FileName = filepath.Base(fPath)
	// Get the file extension
	fileExtension := filepath.Ext(fPath)
	// Remove the leading dot from the extension, if it exists
	if len(fileExtension) > 0 {
		data.FileType = fileExtension[1:]
	} else {
		data.FileType = util.UNKNOWN_FILE_TYPE
	}

	data.File, err = os.ReadFile(fPath)
	if err != nil {
		panic(fmt.Errorf("error while creating a table %v", err))
	}

	return data, sizeOfTheData
}

func getFileSizeGB(filePath string) (float64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	// Get size in bytes and convert to gigabytes
	fileSizeGB := float64(fileInfo.Size()) / (1024 * 1024 * 1024)
	return fileSizeGB, nil
}
