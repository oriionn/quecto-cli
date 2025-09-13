package config

import (
	"os"
)

func IsConfigExists() bool {
	path := GetConfigPath()

	_, err := os.Stat(path)
	return err == nil
}

func CreateConfig() error {
	path := GetConfigDir()
	os.MkdirAll(path, os.ModePerm)

	path = GetConfigPath()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("https://s.oriondev.fr/")
	return nil
}
