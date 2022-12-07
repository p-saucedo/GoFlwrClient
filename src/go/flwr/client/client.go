package goflwr

type IClient interface {
	GetParameters(config map[string]interface{}) [][]byte
	GetProperties(config map[string]interface{}) map[string]interface{}
	Fit([][]byte, map[string]interface{}) ([][]byte, int, map[string]interface{})
	Evaluate([][]byte, map[string]interface{}) (float32, int, map[string]interface{})
}
