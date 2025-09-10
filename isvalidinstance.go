package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/oriionn/quecto-cli/utils"
	"github.com/urfave/cli/v3"
)

func IsValidInstance(ctx context.Context, cmd *cli.Command) error {
	domain := cmd.StringArg("domain")
	url, err := utils.FormatDomain(domain)
	if err != nil {
		return utils.PrintError(err)
	}

	reqUrl := fmt.Sprintf("%sapi/config", url)
	res, err := http.Get(reqUrl)
	if err != nil {
		return utils.PrintError(err)
	}

	if res.StatusCode == 200 {
		msg := fmt.Sprintf("%s is a Quecto instance.", domain)
		fmt.Println(utils.SuccessStyle.Render(msg))
		return nil
	}

	msg := fmt.Sprintf("%s is not a Quecto instance.", domain)
	fmt.Println(utils.ErrorStyle.Render(msg))
	return nil
}
