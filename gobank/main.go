package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	err = store.init()
	if err != nil {
		log.Fatal(err)
	}
	server := NewApiServer(":3000", store)
	server.Run()
}
