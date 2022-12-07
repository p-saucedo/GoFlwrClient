package goflwr

import (
	serde "goflwr/src/go/flwr/common"
	typing "goflwr/src/go/flwr/typing"
)

type ClientWrapper struct {
	Client IClient
}

func (cw *ClientWrapper) GetParameters(ins *typing.GetParametersIns) *typing.GetParametersRes {

	parameters := cw.Client.GetParameters(serde.ConfigToMap(ins.Config))
	return &typing.GetParametersRes{Parameters: *TensorToParameters(parameters)}
}

func (cw *ClientWrapper) GetProperties(ins *typing.GetPropertiesIns) *typing.GetPropertiesRes {
	properties := cw.Client.GetProperties(serde.ConfigToMap(ins.Config))
	return &typing.GetPropertiesRes{Properties: serde.MapToProperties(properties)}
}

func (cw *ClientWrapper) Fit(ins *typing.FitIns) *typing.FitRes {
	parameters, numExamples, metrics := cw.Client.Fit(ParametersToTensor(&ins.Parameters), serde.ConfigToMap(ins.Config))
	return &typing.FitRes{Parameters: *TensorToParameters(parameters), NumExamples: int64(numExamples), Metrics: serde.MapToMetrics(metrics)}
}

func (cw *ClientWrapper) Evaluate(ins *typing.EvaluateIns) *typing.EvaluateRes {
	loss, numExamples, metrics := cw.Client.Evaluate(ParametersToTensor(&ins.Parameters), serde.ConfigToMap(ins.Config))
	return &typing.EvaluateRes{Loss: loss, NumExamples: int64(numExamples), Metrics: serde.MapToMetrics(metrics)}

}
