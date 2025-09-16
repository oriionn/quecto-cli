package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/oriionn/quecto-cli/v2/utils"
	"github.com/urfave/cli/v3"
)

type CheckPasswordResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data bool `json:"data"`
}

type UnshortenResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data UnshortenResponseData `json:"data"`
}

type UnshortenResponseData struct {
	ShortCode string `json:"short_code"`
	Link string `json:"link"`
	Expiration int `json:"expiration"`
	CreatedAt string `json:"created_at"`
}

func Unshorten(ctx context.Context, cmd *cli.Command) error {
	link := cmd.StringArg("link")
	password := cmd.String("password")

	if utils.IsEmpty(link) {
		return utils.PrintErr("You haven't specified a link to unshorten.")
	}

	if !utils.UrlRegex.MatchString(link) {
		return utils.PrintErr("Invalid URL")
	}

	parsed, err := url.Parse(link)
	if err != nil {
		return utils.PrintError(err)
	}

	domain := parsed.Host
	short_code := strings.Trim(parsed.Path, "/")

	url, err := utils.FormatDomain(domain)
	if err != nil {
		return utils.PrintError(err)
	}

	isValid, err := utils.IsValidInstance(url)
	if err != nil {
		return utils.PrintError(err)
	}
	if !isValid {
		return utils.PrintErr("Invalid link.")
	}

	fetchUrl := fmt.Sprintf("%sapi/%s/password", url, short_code)
	res, err := http.Get(fetchUrl)
	if err != nil {
		return utils.PrintError(err)
	}
	defer res.Body.Close()

	var result CheckPasswordResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return utils.PrintError(err)
	}

	if res.StatusCode != 200 {
		if !utils.IsEmpty(result.Message) {
			return utils.PrintErr(result.Message)
		}

		return utils.PrintErr("An internal error occurred on the instance.")
	}

	if result.Data && utils.IsEmpty(password) {
		return utils.PrintErr("The link requires a password")
	}

	fetchUrl = fmt.Sprintf("%sapi/%s", url, short_code)
	res, err = http.Get(fetchUrl)
	if err != nil {
		return utils.PrintError(err)
	}
	defer res.Body.Close()

	var unshortenResult UnshortenResponse
	err = json.NewDecoder(res.Body).Decode(&unshortenResult)
	if err != nil {
		return utils.PrintError(err)
	}

	if res.StatusCode != 200 {
		if !utils.IsEmpty(result.Message) {
			return utils.PrintErr(result.Message)
		}

		return utils.PrintErr("An internal error occurred on the instance.")
	}

	fmt.Println(utils.SuccessStyle.Render(unshortenResult.Data.Link))
    return nil
}
