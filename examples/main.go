package main

import (
	"log"

	"github.com/parikshitg/gomongo"
)

func main() {

	uri := "mongodb://localhost:27017"
	client, err := gomongo.Connect(uri)
	if err != nil {
		log.Fatal("Connection Error : ", err)
	}
	log.Println("client : ", client)
}
