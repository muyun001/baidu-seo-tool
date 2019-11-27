package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// 获取主域
func Domain(originalUrl string) (domainHost string, err error) {
	originalUrl = strings.Replace(originalUrl, "http://", "", -1)
	originalUrl = strings.Replace(originalUrl, "https://", "", -1)
	if !strings.Contains(originalUrl, "/") && strings.Contains(originalUrl, "...") {
		return "", nil
	}

	originalUrl = strings.Replace(originalUrl, "...", "", -1)
	originalUrl = FormatUrl(originalUrl)
	domain, err := url.Parse(originalUrl)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return domain.Host, nil
}

// 检查scheme，没有则加上http://
func FormatUrl(url string) string {
	url = strings.Trim(url, " ")
	schemeRegx := regexp.MustCompile(`^https?://.*`)
	if false == schemeRegx.MatchString(url) {
		url = "http://" + url
	}

	return url
}

