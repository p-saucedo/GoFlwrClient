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

func (cw *ClientWrapper) Fit() {

}

func (cw *ClientWrapper) Evaluate() {

}
