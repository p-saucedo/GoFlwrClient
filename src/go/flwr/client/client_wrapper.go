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
	return &typing.GetParametersRes{Parameters: typing.Parameters{Tensors: parameters, TensorType: "blabla"}}
}

func (cw *ClientWrapper) GetProperties(ins *typing.GetPropertiesIns) *typing.GetPropertiesRes {
	properties := cw.Client.GetProperties(serde.ConfigToMap(ins.Config))
	return &typing.GetPropertiesRes{Properties: serde.MapToProperties(properties)}
}

func (cw *ClientWrapper) Fit(ins *typing.FitIns) *typing.FitRes {
	parameters, numExamples, metrics := cw.Client.Fit(ins.Parameters.Tensors, serde.ConfigToMap(ins.Config))
	return &typing.FitRes{Parameters: typing.Parameters{Tensors: parameters, TensorType: "kk"}, NumExamples: int64(numExamples), Metrics: serde.MapToMetrics(metrics)}
}

func (cw *ClientWrapper) Evaluate(ins *typing.EvaluateIns) *typing.EvaluateRes {
	loss, numExamples, metrics := cw.Client.Evaluate(ins.Parameters.Tensors, serde.ConfigToMap(ins.Config))
	return &typing.EvaluateRes{Loss: loss, NumExamples: int64(numExamples), Metrics: serde.MapToMetrics(metrics)}

}
