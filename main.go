package main

import (
	c "goflwr/src/go/flwr/client"
)

func main() {
	client := &c.CustomClient{}

	c.StartClient("127.0.0.1:8080", client)

}
