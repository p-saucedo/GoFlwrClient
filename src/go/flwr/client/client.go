package goflwr

import (
	tensor "gorgonia.org/tensor"
)

type IClient interface {
	GetParameters(config map[string]interface{}) []*tensor.Dense
	GetProperties(config map[string]interface{}) map[string]interface{}
	Fit([]*tensor.Dense, map[string]interface{}) ([]*tensor.Dense, int, map[string]interface{})
	Evaluate([]*tensor.Dense, map[string]interface{}) (float32, int, map[string]interface{})
}
