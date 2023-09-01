package main

import (
	"flag"
	"fmt"
	"littleapi/api"
	"littleapi/storage"
	"log"
)

func main() {
	port := flag.String("port", ":3000", "Port to listen on")
	flag.Parse()

	store := storage.NewDbStorage()

	server := api.NewServer(*port, store)
	fmt.Println("server is running on port: ", *port)
	log.Fatal(server.Start())
}
