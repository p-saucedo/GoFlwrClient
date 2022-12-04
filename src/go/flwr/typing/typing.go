package goflwr

type Scalar interface{} // bool, bytes, float64, int, string
type Config map[string]Scalar
type Properties map[string]Scalar

type Parameters struct {
	Tensors    [][]byte
	TensorType string
}

type GetParametersIns struct {
	Config Config
}

type GetParametersRes struct {
	Config     Config
	Parameters Parameters
}
