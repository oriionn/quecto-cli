package main

import (
	"context"
	"fmt"

	"github.com/oriionn/quecto-cli/config"
	"github.com/oriionn/quecto-cli/utils"
	"github.com/urfave/cli/v3"
)

func ConfigGet(ctx context.Context, cmd *cli.Command) error {
	config, err := config.Get()
	if err != nil {
		return utils.PrintError(err)
	}

	fmt.Println(utils.SuccessStyle.Render(config))
	return nil
}

func ConfigSet(ctx context.Context, cmd *cli.Command) error {
	newDomain := cmd.StringArg("domain")
	if utils.IsEmpty(newDomain) {
		return utils.PrintErr("You haven't specified a domain to set as your default instance.")
	}

	url, err := utils.FormatDomain(newDomain)
	if err != nil {
		return utils.PrintError(err)
	}

	isValid, err := utils.IsValidInstance(url)
	if !isValid {
		return utils.PrintErr("The instance is invalid.")
	}

	err = config.Set(newDomain)
	if err != nil {
		return utils.PrintError(err)
	}
	msg := fmt.Sprintf("Your default instance domain has been set to `%s`.", newDomain)
	fmt.Println(utils.SuccessStyle.Render(msg))

	return nil
}
