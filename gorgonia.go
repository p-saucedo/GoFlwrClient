package main

import (
	"flag"
	"fmt"
	"log"

	_ "net/http/pprof"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

var (
	epochs     = flag.Int("epochs", 10, "Number of epochs to train for")
	dataset    = flag.String("dataset", "train", "Which dataset to train on? Valid options are \"train\" or \"test\"")
	dtype      = flag.String("dtype", "float64", "Which dtype to use")
	batchsize  = flag.Int("batchsize", 10, "Batch size")
	cpuprofile = flag.String("cpuprofile", "", "CPU profiling")
)

const loc = "./mnist/"

var dt tensor.Dtype

func parseDtype() {
	switch *dtype {
	case "float64":
		dt = tensor.Float64
	case "float32":
		dt = tensor.Float32
	default:
		log.Fatalf("Unknown dtype: %v", *dtype)
	}
}

type nn struct {
	g      *gorgonia.ExprGraph
	w0, w1 *gorgonia.Node

	pred    *gorgonia.Node
	predVal gorgonia.Value
}

func newNN(g *gorgonia.ExprGraph) *nn {
	// Create node for w/weight
	w0 := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(2, 2), gorgonia.WithName("w0"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))
	w1 := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(2, 1), gorgonia.WithName("w1"), gorgonia.WithInit(gorgonia.GlorotU(1.0)))
	return &nn{
		g:  g,
		w0: w0,
		w1: w1,
	}
}

func (m *nn) learnables() gorgonia.Nodes {
	return gorgonia.Nodes{m.w0, m.w1}
}

func (m *nn) fwd(x *gorgonia.Node) (err error) {
	var l0, l1, l2 *gorgonia.Node
	var l0dot, l1dot *gorgonia.Node

	// Camada de input
	l0 = x

	// Multiplicação pelos pesos e sigmoid
	l0dot = gorgonia.Must(gorgonia.Mul(l0, m.w0))

	// Input para a hidden layer
	l1 = gorgonia.Must(gorgonia.Sigmoid(l0dot))

	// Multiplicação pelos pesos:
	l1dot = gorgonia.Must(gorgonia.Mul(l1, m.w1))

	// Camada de saída:
	l2 = gorgonia.Must(gorgonia.Sigmoid(l1dot))

	m.pred = l2
	gorgonia.Read(m.pred, &m.predVal)
	return nil

}

func main() {
	parseDtype()
	g := gorgonia.NewGraph()
	m := newNN(g)

	xB := []float64{1, 0, 0, 1, 1, 1, 0, 0}
	xT := tensor.New(tensor.WithBacking(xB), tensor.WithShape(4, 2))
	x := gorgonia.NewMatrix(g,
		tensor.Float64,
		gorgonia.WithName("X"),
		gorgonia.WithShape(4, 2),
		gorgonia.WithValue(xT),
	)

	// Define validation data set
	yB := []float64{1, 1, 0, 0}
	yT := tensor.New(tensor.WithBacking(yB), tensor.WithShape(4, 1))
	y := gorgonia.NewMatrix(g,
		tensor.Float64,
		gorgonia.WithName("y"),
		gorgonia.WithShape(4, 1),
		gorgonia.WithValue(yT),
	)

	if err := m.fwd(x); err != nil {
		log.Fatalf("%+v", err)
	}

	losses := gorgonia.Must(gorgonia.Sub(y, m.pred))
	square := gorgonia.Must(gorgonia.Square(losses))
	cost := gorgonia.Must(gorgonia.Mean(square))

	if _, err := gorgonia.Grad(cost, m.learnables()...); err != nil {
		log.Fatal(err)
	}

	vm := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(m.learnables()...))
	solver := gorgonia.NewVanillaSolver(gorgonia.WithLearnRate(0.1))

	for i := 0; i < 10000; i++ {
		vm.Reset()
		if err := vm.RunAll(); err != nil {
			log.Fatalf("Failed at inter  %d: %v", i, err)
		}
		solver.Step(gorgonia.NodesToValueGrads(m.learnables()))
		vm.Reset()
	}
	fmt.Println("\n\nOutput after Training: \n", m.predVal)

}
