package config

import (
	"os"
)

func Get() (string, error) {
	if !IsConfigExists() {
		CreateConfig()
	}

	path := GetConfigPath()
	domain, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(domain), nil
}
