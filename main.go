package main

import (
	"github.com/polivera/proxy-service/src/data"
	"github.com/polivera/proxy-service/src/proxy"
	"log"
)

func main() {
	var (
		err   error
		store data.Store
	)

	// All this should come from config
	dbDriver := "sqlite"
	dbUrl := "./test.db"
	proxyHost := "localhost:5000"

	if store, err = data.NewStore(dbDriver, dbUrl); err != nil {
		log.Fatalf("Error connecting to the database. Error: %s", err)
	}
	if err = store.Migrate(); err != nil {
		log.Fatalf("Error initializing / migrating tables. Error: %s", err)
	}

	proxyServer := proxy.GetProxy(proxyHost, store)
	proxyServer.Run()
}
