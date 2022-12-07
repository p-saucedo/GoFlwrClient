package main

import (
	c "goflwr/src/go/flwr/client"
	"log"
)

func main() {
	client := &c.CustomClient{}
	log.Println("Starting client...")
	err := c.StartClient("127.0.0.1:8080", client)
	if err != nil {
		log.Print(err)
	}

}
