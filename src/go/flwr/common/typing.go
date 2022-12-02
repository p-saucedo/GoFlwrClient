package goflwr

type Scalar interface{} // bool, bytes, float64, int, string
type Config map[string]Scalar
type Properties map[string]Scalar

type Parameters struct {
	tensors     []byte
	tensor_type string
}

type GetParametersIns struct {
	config Config
}

type GetParametersRes struct {
	config     Config
	parameters Parameters
}
