package config

import "os"

func Set(domain string) error {
	if !IsConfigExists() {
		err := CreateConfig()

		if err != nil {
			return err
		}
	}

	path := GetConfigPath()
	return os.WriteFile(path, []byte(domain), os.ModePerm)
}
