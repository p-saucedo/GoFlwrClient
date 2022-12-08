package goflwr

// https://github.com/gorgonia/tensor/blob/master/dense_io.go
import (
	tensor "gorgonia.org/tensor"
)

type CustomClient struct {
	IClient
}

func (client *CustomClient) GetParameters(config map[string]interface{}) []*tensor.Dense {
	m := tensor.New(tensor.WithShape(2, 2), tensor.WithBacking([]float64{1, 2, 3, 4}))

	var p = [1]*tensor.Dense{m}

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
