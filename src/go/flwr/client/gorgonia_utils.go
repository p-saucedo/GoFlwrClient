package goflwr

import (
	G "gorgonia.org/gorgonia"
	tensor "gorgonia.org/tensor"
)

func GorgoniaGetDataOfNode(n *G.Node) *tensor.Dense {
	return tensor.NewDense(n.Dtype(), n.Shape(), tensor.WithBacking(n.Value().Data()))
}

func GorgoniaSetDataOfNode(d *tensor.Dense, n *G.Node) error {
	return G.Let(n, d)
}
