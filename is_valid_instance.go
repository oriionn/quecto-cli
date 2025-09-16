package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/oriionn/quecto-cli/v2/utils"
	"github.com/urfave/cli/v3"
)

func IsValidInstance(ctx context.Context, cmd *cli.Command) error {
	domain := cmd.StringArg("domain")
	if strings.TrimSpace(domain) == "" {
		return utils.PrintErr("You haven't specified a domain.")
	}

	url, err := utils.FormatDomain(domain)
	if err != nil {
		return utils.PrintError(err)
	}

	isValid, err := utils.IsValidInstance(url)
	if err != nil {
		return utils.PrintError(err)
	}

	if isValid {
		msg := fmt.Sprintf("%s is a valid Quecto instance.", domain)
		fmt.Println(utils.SuccessStyle.Render(msg))
	} else {
		msg := fmt.Sprintf("%s is a invalid Quecto instance.", domain)
		fmt.Println(utils.ErrorStyle.Render(msg))
	}

	return nil
}
