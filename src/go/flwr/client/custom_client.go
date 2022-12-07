package goflwr

type CustomClient struct {
	IClient
}

func (client *CustomClient) GetParameters(config map[string]interface{}) [][]byte {
	return make([][]byte, 4)
}

func (client *CustomClient) GetProperties(config map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"hola": 1}
}

func (client *CustomClient) Fit([][]byte, map[string]interface{}) ([][]byte, int, map[string]interface{}) {
	return make([][]byte, 4), 18, map[string]interface{}{"test": "fit"}
}

func (client *CustomClient) Evaluate([][]byte, map[string]interface{}) (float32, int, map[string]interface{}) {
	return 0.12, 8, map[string]interface{}{"test": "evaluate"}
}
