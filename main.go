package main

import (
	_ "go-videosite-server/conf"
	"go-videosite-server/routers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := routers.NewRouter()

	httpServer := http.Server{
		Addr:    os.Getenv("ADDR"),
		Handler: router,
	}

	log.Printf("Server started on %s", os.Getenv("ADDR"))
	err := httpServer.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
