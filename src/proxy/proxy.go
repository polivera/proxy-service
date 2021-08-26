package proxy

import (
	"fmt"
	"github.com/polivera/proxy-service/src/data"
	"github.com/polivera/proxy-service/src/data/models"
	"github.com/polivera/proxy-service/src/utils"
	"log"
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
	// proxy the request
	// store the response IF the response is successful
	// return proxy request

	var (
		config models.RequestConfig
		err    error
	)
	if config, err = prx.getConfig(r); err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Config for request does not exist"))
	}

	fmt.Println(config)
	w.Write([]byte("Shit working yo"))
}

func (prx *proxy) saveRequest(w http.ResponseWriter, r *http.Request) {

}

func (prx *proxy) getConfig(r *http.Request) (models.RequestConfig, error) {
	host, _ := utils.GetFullURLFromRequest(r)
	return prx.db.GetConfig(host)

}
