package main

import (
	"log"
	"my-app/functional_option_pattern/server"
	"os"
	"time"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	svr := server.New("localhost", 8888,
		server.WithTimeout(time.Minute),
		server.WithLogger(logger))
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
