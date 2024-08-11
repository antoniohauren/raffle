package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antoniohauren/raffle/config"
	"github.com/antoniohauren/raffle/database"
	"github.com/antoniohauren/raffle/router"
)

func init() {
	config.Init()
}

func main() {
	conn := database.ConnectDatabase()
	server := http.NewServeMux()

	// ping router
	server.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	// routes
	server.Handle("/raffle/", http.StripPrefix("/raffle", router.NewRaffleRouter(conn)))

	// start
	entry := fmt.Sprintf("localhost:%s", config.Port)

	log.Println("running on " + entry)

	if err := http.ListenAndServe(entry, server); err != nil {
		log.Fatal("something gone wrong")
	}
}
