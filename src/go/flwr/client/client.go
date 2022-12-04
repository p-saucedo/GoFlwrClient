package goflwr

type IClient interface {
	GetParameters(config map[string]interface{}) [][]byte
	GetProperties(config map[string]interface{}) map[string]interface{}
	Fit()
	Evaluate()
}
