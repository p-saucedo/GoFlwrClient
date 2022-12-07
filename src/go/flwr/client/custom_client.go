package goflwr

// https://github.com/gorgonia/tensor/blob/master/dense_io.go
import (
	"bytes"
	"log"

	tensor "gorgonia.org/tensor"
)

type CustomClient struct {
	IClient
}

func (client *CustomClient) GetParameters(config map[string]interface{}) [][]byte {
	m := tensor.New(tensor.WithShape(2, 2), tensor.WithBacking([]float64{1, 2, 3, 4}))

	buf := new(bytes.Buffer)
	err := m.WriteNpy(buf)

	if err != nil {
		panic(err)
	}

	log.Print(buf)

	return make([][]byte, 4)
}

func (client *CustomClient) GetProperties(config map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"hola": 1}
}

func (client *CustomClient) Fit([][]byte, map[string]interface{}) ([][]byte, int, map[string]interface{}) {
	return make([][]byte, 4), 12, map[string]interface{}{"test": "fit"}
}

func (client *CustomClient) Evaluate([][]byte, map[string]interface{}) (float32, int, map[string]interface{}) {
	return 0.12, 2, map[string]interface{}{"test": "evaluate"}
}
