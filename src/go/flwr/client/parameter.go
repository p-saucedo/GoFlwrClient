package goflwr

// https://github.com/gorgonia/tensor/blob/master/dense_io.go
import (
	"bytes"
	typing "goflwr/src/go/flwr/typing"

	tensor "gorgonia.org/tensor"
)

func TensorToParameters(t *tensor.Dense) *typing.Parameters {
	buf := new(bytes.Buffer)
	err := t.WriteNpy(buf)

	if err != nil {
		panic(err)
	}

	return &typing.Parameters{Tensors: buf.Bytes(), TensorType: "numpy.ndarray"}
}

func ParametersToTensor(p *typing.Parameters) *tensor.Dense {
	var m *tensor.Dense

	err := m.ReadNpy(bytes.NewReader(p.Tensors))

	if err != nil {
		panic(err)
	}

	return m
}
