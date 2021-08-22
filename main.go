package main

import (
	"fmt"
	"github.com/polivera/proxy-service/src/data"
	"github.com/polivera/proxy-service/src/proxy"
	"log"
)

func main() {
	var (
		err   error
		store data.Store
	)
	if store, err = data.NewStore("sqlite", "./test.db"); err != nil {
		log.Fatalf("Error connecting to the database. Error: %s", err)
	}
	if err = store.Migrate(); err != nil {
		log.Fatalf("Error initializing / migrating tables. Error: %s", err)
	}

	proxyServer := proxy.GetProxy("localhost:5000", store)
	proxyServer.Run()

	fmt.Println("This is working")
}
