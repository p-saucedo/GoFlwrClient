package goflwr

type Scalar interface{} // bool, bytes, float64, int64, string
type Config map[string]Scalar
type Metrics map[string]Scalar
type Properties map[string]Scalar

type Parameters struct {
	Tensors    []byte
	TensorType string
}

type GetParametersIns struct {
	Config Config
}

type GetParametersRes struct {
	Config     Config
	Parameters Parameters
}

type GetPropertiesIns struct {
	Config Config
}

type GetPropertiesRes struct {
	Config     Config
	Properties Properties
}

type FitIns struct {
	Parameters Parameters
	Config     Config
}

type FitRes struct {
	Parameters  Parameters
	Metrics     Metrics
	NumExamples int64
}

type EvaluateIns struct {
	Parameters Parameters
	Config     Config
}

type EvaluateRes struct {
	Loss        float32
	Metrics     Metrics
	NumExamples int64
}
