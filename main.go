package main

import (
	c "goflwr/src/go/flwr/client"
	"log"

	G "gorgonia.org/gorgonia"
	tensor "gorgonia.org/tensor"
)

type CustomClient struct {
	c.IClient
}

func (client *CustomClient) GetParameters(config map[string]interface{}) []*tensor.Dense {
	return []*tensor.Dense{getDataOfNode(NN.w0), getDataOfNode(NN.w1)}
}

func (client *CustomClient) GetProperties(config map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"num_pesos": 2}
}

func (client *CustomClient) Fit(parameters []*tensor.Dense, config map[string]interface{}) ([]*tensor.Dense, int, map[string]interface{}) {
	var l0, l1, l2 *G.Node
	var l0dot, l1dot *G.Node

	// FWD
	l0 = x
	l0dot = G.Must(G.Mul(l0, NN.w0))
	l1 = G.Must(G.Sigmoid(l0dot))
	l1dot = G.Must(G.Mul(l1, NN.w1))
	l2 = G.Must(G.Sigmoid(l1dot))
	NN.pred = l2
	G.Read(NN.pred, &NN.predVal)

	//BCKWD
	losses := G.Must(G.Sub(y, NN.pred))
	square := G.Must(G.Square(losses))
	cost := G.Must(G.Mean(square))

	if _, err := G.Grad(cost, NN.learnables()...); err != nil {
		log.Fatal(err)
	}

	vm := G.NewTapeMachine(g, G.BindDualValues(NN.learnables()...))
	solver := G.NewVanillaSolver(G.WithLearnRate(0.1))

	for i := 0; i < 1; i++ {
		vm.Reset()
		if err := vm.RunAll(); err != nil {
			log.Fatalf("Failed at inter  %d: %v", i, err)
		}
		solver.Step(G.NodesToValueGrads(NN.learnables()))
		vm.Reset()
	}

	return []*tensor.Dense{getDataOfNode(NN.w0), getDataOfNode(NN.w1)}, x.Shape()[0], map[string]interface{}{"metrics": "fit"}
}

func (client *CustomClient) Evaluate(parameters []*tensor.Dense, config map[string]interface{}) (float32, int, map[string]interface{}) {
	return 0.12, 2, map[string]interface{}{"test": "evaluate"}
}

// --------------------------------------------------------------------
type nn struct {
	g      *G.ExprGraph
	w0, w1 *G.Node

	pred    *G.Node
	predVal G.Value
}

func getDataOfNode(n *G.Node) *tensor.Dense {

	return tensor.NewDense(n.Dtype(), n.Shape(), tensor.WithBacking(n.Value().Data()))
}

func newNN(g *G.ExprGraph) *nn {
	// Create node for w/weight
	w0 := G.NewMatrix(g, tensor.Float64, G.WithShape(2, 2), G.WithName("w0"), G.WithInit(G.GlorotN(1.0)))
	w1 := G.NewMatrix(g, tensor.Float64, G.WithShape(2, 1), G.WithName("w1"), G.WithInit(G.GlorotN(1.0)))
	return &nn{
		g:  g,
		w0: w0,
		w1: w1,
	}
}

func (m *nn) learnables() G.Nodes {
	return G.Nodes{m.w0, m.w1}
}

var g = G.NewGraph()
var NN = newNN(g)

var xB = []float64{1, 0, 0, 1, 1, 1, 0, 0}
var xT = tensor.New(tensor.WithBacking(xB), tensor.WithShape(4, 2))
var x = G.NewMatrix(g,
	tensor.Float64,
	G.WithName("X"),
	G.WithShape(4, 2),
	G.WithValue(xT),
)

// Define validation data set
var yB = []float64{1, 1, 0, 0}
var yT = tensor.New(tensor.WithBacking(yB), tensor.WithShape(4, 1))
var y = G.NewMatrix(g,
	tensor.Float64,
	G.WithName("y"),
	G.WithShape(4, 1),
	G.WithValue(yT),
)

// -------------------------------------------------------------------

func main() {

	client := &CustomClient{}
	log.Println("Starting client...")
	err := c.StartClient("127.0.0.1:8080", client)
	if err != nil {
		log.Print(err)
	}

}
