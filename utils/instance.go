package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ConfigResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data ConfigResponseData `json:"data"`
}

type ConfigResponseData struct {
	Instance string `json:"instance"`
	Domain string `json:"domain"`
	SSL bool `json:"ssl"`
	AuthorizeCustomCode bool `json:"authorize_custom_shortcode"`
	Expirations []ConfigResponseExpiration `json:"expirations"`
}

type ConfigResponseExpiration struct {
	Name string `json:"name"`
	Minutes int `json:"minutes"`
}

func IsValidInstance(url string) (bool, error) {
	reqUrl := fmt.Sprintf("%sapi/config", url)
	res, err := http.Get(reqUrl)
	if err != nil {
		return false, err
	}

	if res.StatusCode == 200 {
		return true, nil
	}

	return false, nil
}

func GetConfig(url string) (ConfigResponseData, error) {
	reqUrl := fmt.Sprintf("%sapi/config", url)
	res, err := http.Get(reqUrl)
	if err != nil {
		return ConfigResponseData{}, err
	}

	if res.StatusCode != 200 {
		return ConfigResponseData{}, fmt.Errorf("Your instance is not valid.")
	}

	var result ConfigResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return ConfigResponseData{}, err
	}

	return result.Data, nil
}
