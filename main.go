package main

import (
	c "goflwr/src/go/flwr/client"
)

func main() {
	r := c.Connect()

	client := c.Client{}

	server_msg, _ := r.Recv()

	c.Handle(client, *server_msg)

}
