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

func (cw *ClientWrapper) GetProperties() {

}

func (cw *ClientWrapper) Fit(ins *typing.FitIns) *typing.FitRes {
	parameters, numExamples, metrics := cw.Client.Fit(ins.Parameters.Tensors, serde.ConfigToMap(ins.Config))
	return &typing.FitRes{Parameters: typing.Parameters{Tensors: parameters, TensorType: "kk"}, NumExamples: numExamples, Metrics: serde.MetricsToMap()}
}

func (cw *ClientWrapper) Evaluate() {

}
