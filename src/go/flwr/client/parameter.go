package goflwr

// https://github.com/gorgonia/tensor/blob/master/dense_io.go
import (
	"bytes"
	typing "goflwr/src/go/flwr/typing"
	"log"

	tensor "gorgonia.org/tensor"
)

func TensorToParameters(t []*tensor.Dense) *typing.Parameters {

	tensors := make([][]byte, len(t))

	for idx, _tensor := range t {
		tensors[idx] = TensorToBytes(_tensor)
	}

	return &typing.Parameters{Tensors: tensors, TensorType: "numpy.ndarray"}
}

func TensorToBytes(t *tensor.Dense) []byte {
	buf := new(bytes.Buffer)
	err := t.WriteNpy(buf)

	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func ParametersToTensor(p *typing.Parameters) []*tensor.Dense {
	tensors := make([]*tensor.Dense, len(p.Tensors))

	for idx, _tensor := range p.Tensors {
		tensors[idx] = BytesToTensor(_tensor)
	}

	return tensors
}

func BytesToTensor(b []byte) *tensor.Dense {
	m := tensor.NewDense(tensor.Float64, tensor.Shape{1, 1})

	err := m.ReadNpy(bytes.NewReader(b))

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return m

}
