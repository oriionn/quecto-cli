package utils

import (
	"fmt"
	"net/http"
)

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
