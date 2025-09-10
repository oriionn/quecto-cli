package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var urlRegex = regexp.MustCompile(`^https?:\/\/([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(:[0-9]+)?(\/.*)?$`)

func FormatDomain(domain string) (string, error) {
	originalDomain := domain
	if !strings.HasPrefix(domain, "http") {
		domain = fmt.Sprintf("%s%s", "https://", domain)
	}

	if !strings.HasSuffix(domain, "/") {
		domain += "/"
	}

	if !urlRegex.MatchString(domain) {
		return "", fmt.Errorf("invalid domain format: %s", originalDomain)
	}

	return domain, nil
}
