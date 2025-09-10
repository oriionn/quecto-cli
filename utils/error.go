package utils

import "fmt"

func PrintError(err error) error {
	fmt.Println(ErrorStyle.Render(err.Error()))
	return nil
}
