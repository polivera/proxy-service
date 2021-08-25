package utils

import (
	"errors"
	"fmt"
	"github.com/polivera/proxy-service/src"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GetFullURLFromRequest GetFullURL Return URL, URI and full URL from the url.URL object
func GetFullURLFromRequest(r *http.Request) (fullHost string, fullPath string) {
	urlScheme := r.URL.Scheme
	if urlScheme == "" {
		urlScheme = "http"
	}

	urlHost := r.Host
	if urlHost == "" {
		urlHost = r.URL.Host
	}
	fullHost = fmt.Sprintf("%s://%s", urlScheme, urlHost)

	fullPath = r.URL.Path
	if r.URL.RawQuery != "" {
		fullPath += "?" + r.URL.RawQuery
	}

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
