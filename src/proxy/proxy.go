package proxy

import (
	"fmt"
	"github.com/polivera/proxy-service/src/data"
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
	prx.getConfig(r)
	w.Write([]byte("Shit working yo"))
}

func (prx *proxy) saveRequest(w http.ResponseWriter, r *http.Request) {

}

func (prx *proxy) getConfig(r *http.Request) {
	host, path := utils.GetFullURLFromRequest(r)
	conf, err := prx.db.GetConfig(host)

	fmt.Println(host)
	fmt.Println(path)
	fmt.Println(conf)
	fmt.Println(err)
}
