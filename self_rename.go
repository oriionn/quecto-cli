package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/oriionn/quecto-cli/v2/utils"
	"github.com/urfave/cli/v3"
)

func SelfRename(ctx context.Context, cmd *cli.Command) error {
	execPath, err := os.Executable()
	if err != nil {
		return utils.PrintError(err)
	}

	directory := filepath.Dir(execPath)
	newPath := filepath.Join(directory, "quecto")
	err = os.Rename(execPath, newPath)
	if err != nil {
		return utils.PrintError(err)
	}

	fmt.Println(utils.SuccessStyle.Render("The binary has been renamed as `quecto`, to test this, execute `quecto` in your terminal."))
	return nil
}
