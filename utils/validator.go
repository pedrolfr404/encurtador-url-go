package utils

import (
	"net/url"
	"strings"
)

var blockedDomains = []string{
	"bit.ly",
	"t.co",
	"tinyurl.com",
	"cutt.ly",
}

func IsValid(inputURL string) bool {
	u, err := url.ParseRequestURI(inputURL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return false
	}
	return true
}

func IsAlreadyShortened(inputURL string) bool {
	u, err := url.Parse(inputURL)
	if err != nil {
		return false
	}

	host := strings.ToLower(u.Host)
	for _, domain := range blockedDomains {
		if strings.Contains(host, domain) {
			return true
		}
	}
	return false
}
