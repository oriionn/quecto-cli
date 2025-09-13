package config

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetConfigDir() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "quecto-cli")
	}

	return filepath.Join(os.Getenv("HOME"), ".config", "quecto-cli")
}

func GetConfigPath() string {
	path := GetConfigDir()
	return filepath.Join(path, ".domain")
}
