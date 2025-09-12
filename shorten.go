package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/oriionn/quecto-cli/utils"
	"github.com/urfave/cli/v3"
)

type ShortenBody struct {
	Link            string `json:"link"`
	Expiration      int    `json:"expiration"`
	Password        string `json:"password"`
	CustomShortCode string `json:"custom_sc"`
}

type ShortenResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    ShortenResponseData `json:"data"`
}

type ShortenResponseData struct {
	ShortCode  string `json:"short_code"`
	Link       string `json:"link"`
	Expiration int    `json:"expiration"`
	CreatedAt  string `json:"created_at"`
}

func Shorten(ctx context.Context, cmd *cli.Command) error {
	domain := cmd.String("domain")
	expiration := cmd.Int("expires")
	password := cmd.String("password")
	shortCode := cmd.String("code")
	link := cmd.StringArg("link")

	if utils.IsEmpty(domain) {
		domain = "https://s.oriondev.fr/"
	}

	if utils.IsEmpty(link) {
		fmt.Println(utils.ErrorStyle.Render("You haven't specified a link to shorten."))
		return nil
	}

	url, err := utils.FormatDomain(domain)
	if err != nil {
		return utils.PrintError(err)
	}

	isValid, err := utils.IsValidInstance(url)
	if err != nil {
		return utils.PrintError(err)
	}

	if !isValid {
		fmt.Println(utils.ErrorStyle.Render("Your instance is invalid"))
		return nil
	}

	if !utils.UrlRegex.MatchString(link) {
		fmt.Println(utils.ErrorStyle.Render("Invalid URL"))
		return nil
	}

	body := ShortenBody{
		Link:       link,
		Expiration: expiration,
	}

	if !utils.IsEmpty(password) {
		body.Password = password
	}
	if !utils.IsEmpty(shortCode) {
		body.CustomShortCode = shortCode
	}

	bodyString, err := json.Marshal(body)
	if err != nil {
		return utils.PrintError(err)
	}

	fetchUrl := fmt.Sprintf("%sapi/shorten", url)
	bodyReader := strings.NewReader(string(bodyString))
	req, err := http.NewRequest(http.MethodPost, fetchUrl, bodyReader)
	if err != nil {
		return utils.PrintError(err)
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return utils.PrintError(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println(utils.ErrorStyle.Render("An internal error occurred on the instance."))
		return nil
	}

	var result ShortenResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return utils.PrintError(err)
	}

	link = fmt.Sprintf("%s%s", url, result.Data.ShortCode)
	fmt.Println(utils.SuccessStyle.Render(link))
	return nil
}
