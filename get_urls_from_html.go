package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
    if baseURL == nil || !baseURL.IsAbs() {
        return nil, fmt.Errorf("couldn't parse base URL")
    }

    doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
    if err != nil {
        return nil, fmt.Errorf("couldn't parse HTML: %w", err)
    }

    var urls []string
    doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
        href, ok := s.Attr("href")
        if !ok {
            return
        }
        href = strings.TrimSpace(href)
        if href == "" {
            return
        }
        u, err := url.Parse(href)
        if err != nil {
            return
        }
        abs := baseURL.ResolveReference(u)
        if !abs.IsAbs() {
            return
        }
        urls = append(urls, abs.String())
    })
    return urls, nil
}
