package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const (
	ConfigPath        = "configPath"
	UNKNOWN_FILE_TYPE = "unkown"
)

func GetDefaultDownloadPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "windows":
		return filepath.Join(homeDir, "Downloads"), nil
	case "darwin":
		return filepath.Join(homeDir, "Downloads"), nil
	case "linux":
		return filepath.Join(homeDir, "Downloads"), nil
	default:
		return "", fmt.Errorf("unsupported OS")
	}
}
