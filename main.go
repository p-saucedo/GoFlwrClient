package main

import (
	c "goflwr/src/go/flwr/client"
	"log"

	tensor "gorgonia.org/tensor"
)

type CustomClient struct {
	c.IClient
}

func (client *CustomClient) GetParameters(config map[string]interface{}) []*tensor.Dense {
	return p[:]
}

func (client *CustomClient) GetProperties(config map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"hola": 1}
}

func (client *CustomClient) Fit(parameters []*tensor.Dense, config map[string]interface{}) ([]*tensor.Dense, int, map[string]interface{}) {
	return parameters, 12, map[string]interface{}{"test": "fit"}
}

func (client *CustomClient) Evaluate(parameters []*tensor.Dense, config map[string]interface{}) (float32, int, map[string]interface{}) {
	return 0.12, 2, map[string]interface{}{"test": "evaluate"}
}

var p = [1]*tensor.Dense{tensor.New(tensor.WithShape(2, 2), tensor.WithBacking([]float64{1, 2, 3, 4}))}

func main() {

	client := &CustomClient{}
	log.Println("Starting client...")
	err := c.StartClient("127.0.0.1:8080", client)
	if err != nil {
		log.Print(err)
	}

}
