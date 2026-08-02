package main

import (
	"log"
	"os"

	"github.com/go-by-oksy/go-morse-converter/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	srv := server.NewServer(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
