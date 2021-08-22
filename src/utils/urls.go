package utils

import (
	"errors"
	"fmt"
	"github.com/polivera/proxy-service/src"
	"net/url"
	"strconv"
	"strings"
)

// GetFullURL Return URL, URI and full URL from a url.URL object
func GetFullURL(URL *url.URL) (url string, uri string, fullUrl string) {
	url = fmt.Sprintf("%s://%s", URL.Scheme, URL.Host)
	if URL.Scheme == "" {
		url = "http" + url
	}
	uri = URL.Path
	if URL.RawQuery != "" {
		uri += "?" + URL.RawQuery
	}
	fullUrl = url + uri
	return
}

// SplitURL Return separated scheme, host and uri from a string url
func SplitURL(URL string) (scheme string, host string, uri string, err error) {
	// Method
	urlParts := strings.Split(URL, "://")
	if len(urlParts) < 2 {
		return "", "", "", errors.New("malformed url")
	}
	scheme = urlParts[0]
	// Host
	urlParts = strings.SplitN(urlParts[1], "/", 2)
	host = urlParts[0]
	// URI
	if len(urlParts) > 1 {
		uri = "/" + urlParts[1]
	}
	return
}

func GetQueryParam(url *url.URL, key string, defaultValue string) string {
	for queryKey, value := range url.Query() {
		if queryKey == key {
			return strings.Join(value, ",")
		}
	}
	return defaultValue
}

func GetPaginationData(url *url.URL) ([2]int, error) {
	var (
		pageParam, rowsParam string
		page, rows           int
		err                  error
	)

	pageParam = GetQueryParam(url, src.PagePaginationQueryParam, "0")
	rowsParam = GetQueryParam(url, src.LimitPaginationQueryParam, src.DefaultRowsPerPage)

	if page, err = strconv.Atoi(pageParam); err != nil {
		return [2]int{}, err
	}
	if rows, err = strconv.Atoi(rowsParam); err != nil {
		return [2]int{}, err
	}

	return [2]int{page, rows}, nil
}
