package goflwr

type CustomClient struct {
	IClient
}

func (client *CustomClient) GetParameters(config map[string]interface{}) [][]byte {
	return make([][]byte, 20)
}

func (client *CustomClient) GetProperties(config map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"hola": 1}
}
func (client *CustomClient) Fit() {

}

func (client *CustomClient) Evaluate() {

}
