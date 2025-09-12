package utils

import "fmt"

func PrintError(err error) error {
	fmt.Println(ErrorStyle.Render(err.Error()))
	return nil
}

func PrintErr(str string) error {
	fmt.Println(ErrorStyle.Render(str))
	return nil
}
