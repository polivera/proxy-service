package proxy

import (
	"fmt"
	"github.com/polivera/proxy-service/src/data"
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

func (prx *proxy) Run() {
	http.HandleFunc("/", prx.handler)
	fmt.Printf("Starting proxy service on %s\n", prx.host)
	if err := http.ListenAndServe(prx.host, nil); err != nil {
		log.Fatalf("Cannot start server. Error: %s", err)
	}
}

func (prx *proxy) handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shit working yo"))
}

func (prx *proxy) saveRequest(w http.ResponseWriter, r *http.Request) {

}
