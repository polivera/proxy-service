package proxy

import (
	"fmt"
	"github.com/polivera/proxy-service/src/data"
	"github.com/polivera/proxy-service/src/data/models"
	"github.com/polivera/proxy-service/src/utils"
	"log"
	"net"
	"net/http"
)

type IProxy interface {
	Run()
}

type proxy struct {
	host string
	db   data.Store
}

func GetProxy(host string, db data.Store) IProxy {
	return &proxy{host: host, db: db}
}

// Run proxy server
func (prx *proxy) Run() {
	http.HandleFunc("/", prx.handler)
	fmt.Printf("Starting proxy service on %s\n", prx.host)
	if err := http.ListenAndServe(prx.host, nil); err != nil {
		log.Fatalf("Cannot start server. Error: %s", err)
	}
}

// handler Handle incoming request
func (prx *proxy) handler(w http.ResponseWriter, r *http.Request) {

	// Check if cached request exist
	// if the request exist and has proper status
	// return the cached request

	// else
	// store the response IF the response is successful
	// return proxy request

	var (
		config models.RequestConfig
		err    error
	)
	if config, err = prx.getConfig(r); err != nil {
		// todo: log error (config does not exist for url)
		w.WriteHeader(404)
		_, _ = w.Write([]byte("Config for request does not exist"))
		return
	}

	response, _ := prx.proxyRequest(r, config)
	if response.StatusCode < 300 {
		go func() {
			_ = prx.storeResponseToDB(response)
		}()
	}
	_ = prx.copyResponseToResponseWriter(response, w)
}

// getConfig Get request configuration
func (prx *proxy) getConfig(r *http.Request) (models.RequestConfig, error) {
	host, _ := utils.GetFullURLFromRequest(r)
	return prx.db.GetConfig(host)

}

// proxyRequest
func (prx *proxy) proxyRequest(r *http.Request, config models.RequestConfig) (*http.Response, error) {
	var (
		scheme, host string
		err          error
	)

	// Set remote host
	if scheme, host, _, err = utils.SplitURL(config.Target); err != nil {
		// todo: log error (cannot split url)
		return nil, err
	}
	r.Host = host
	r.URL.Host = host
	r.URL.Scheme = scheme
	// RequestURI can't be set in client requests
	r.RequestURI = ""

	// Set headers
	if host, _, err = net.SplitHostPort(r.RemoteAddr); err != nil {
		r.Header.Set("X-Forwarded-For", host)
	}
	r.Header.Set("Origin", config.OriginHeader)
	r.Header.Set("Referer", config.RefererHeader)

	// todo: logic to encode / decode content
	// In the meantime remove all accepted encoding
	r.Header.Del("Accept-Encoding")

	return http.DefaultClient.Do(r)
}

// copyResponseToResponseWriter - Copy response data to the response writer
func (prx *proxy) copyResponseToResponseWriter(rs *http.Response, rw http.ResponseWriter) error {
	var (
		body string
		err  error
	)

	if rs.Body != nil {
		if body, err = utils.ReadAndResetResponseBody(rs); err != nil {
			return err
		}
	}

	// Response headers
	for header, values := range rs.Header {
		for _, value := range values {
			rw.Header().Set(header, value)
		}
	}

	// Response status and body
	rw.WriteHeader(rs.StatusCode)
	_, err = rw.Write([]byte(body))
	return err
}

// storeResponseToDB Store the response data to requests table on database
func (prx *proxy) storeResponseToDB(response *http.Response) error {
	fmt.Println(response)
	return nil
}
