package main

import (
	"strings"
	"net/url"
)

func normalizeURL(url string) (string, error) {
	normalizedURL := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://"), "/")
	
	return normalizedURL, nil

}
