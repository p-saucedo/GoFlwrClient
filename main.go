package main

import (
	c "goflwr/src/go/flwr/client"
	utils "goflwr/src/go/flwr/utils"
	"log"

	G "gorgonia.org/gorgonia"
	tensor "gorgonia.org/tensor"
)

type CustomClient struct {
	c.IClient
}

func (client *CustomClient) GetParameters(config map[string]interface{}) []*tensor.Dense {
	return []*tensor.Dense{c.GorgoniaGetDataOfNode(NN.w1), c.GorgoniaGetDataOfNode(NN.b1),
		c.GorgoniaGetDataOfNode(NN.w2), c.GorgoniaGetDataOfNode(NN.b2)}
}

func (client *CustomClient) GetProperties(config map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"num_pesos": 2}
}

func (client *CustomClient) Fit(parameters []*tensor.Dense, config map[string]interface{}) ([]*tensor.Dense, int, map[string]interface{}) {
	xD, yD := utils.GetXYMat()

	x := G.NodeFromAny(g, xD, G.WithName("x"))
	y := G.NodeFromAny(g, yD, G.WithName("y"))

	// Set Federated Weights
	var err error
	err = c.GorgoniaSetDataOfNode(parameters[0], NN.w1)
	if err != nil {
		panic(err)
	}
	err = c.GorgoniaSetDataOfNode(parameters[1], NN.b1)
	if err != nil {
		panic(err)
	}
	err = c.GorgoniaSetDataOfNode(parameters[2], NN.w2)
	if err != nil {
		panic(err)
	}
	err = c.GorgoniaSetDataOfNode(parameters[3], NN.b2)
	if err != nil {
		panic(err)
	}

	// Training
	var l0, l1 *G.Node
	var h1, h2 *G.Node
	//var a1, a2 *G.Node

	// Forward propagation
	l0 = x

	// Layer 1
	h1 = G.Must(G.Mul(l0, NN.w1))
	//a1 = G.Must(G.Add(h1, NN.b1))
	l1 = G.Must(G.Sigmoid(h1))

	// Layer 2
	h2 = G.Must(G.Mul(l1, NN.w2))
	//a2, _ = G.Add(h2, NN.b2)

	// Output
	NN.pred = G.Must(G.SoftMax(h2))
	G.Read(NN.pred, &NN.predVal)

	// Backward calculation
	losses := G.Must(G.HadamardProd(NN.pred, y))
	mean := G.Must(G.Mean(losses))
	cost := G.Must(G.Neg(mean))

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

	var costValue G.Value = cost.Value()
	G.Read(cost, &costValue)
	log.Printf("Cost: %f", costValue)

	// Extraction of Weights and required info (trainSet len and metrics)

	return []*tensor.Dense{c.GorgoniaGetDataOfNode(NN.w1), c.GorgoniaGetDataOfNode(NN.b1),
		c.GorgoniaGetDataOfNode(NN.w2), c.GorgoniaGetDataOfNode(NN.b2)}, x.Shape()[0], map[string]interface{}{"metrics": "fit"}
}

func (client *CustomClient) Evaluate(parameters []*tensor.Dense, config map[string]interface{}) (float32, int, map[string]interface{}) {

	xD, yD := utils.GetXYMat()

	x := G.NodeFromAny(g, xD, G.WithName("x"))
	y := G.NodeFromAny(g, yD, G.WithName("y"))

	// Set Federated Weights
	var err error
	err = c.GorgoniaSetDataOfNode(parameters[0], NN.w1)
	if err != nil {
		panic(err)
	}
	err = c.GorgoniaSetDataOfNode(parameters[1], NN.b1)
	if err != nil {
		panic(err)
	}
	err = c.GorgoniaSetDataOfNode(parameters[2], NN.w2)
	if err != nil {
		panic(err)
	}
	err = c.GorgoniaSetDataOfNode(parameters[3], NN.b2)
	if err != nil {
		panic(err)
	}

	// Training
	var l0, l1 *G.Node
	var h1, h2 *G.Node
	//var a1, a2 *G.Node

	// Forward propagation
	l0 = x

	// Layer 1
	h1 = G.Must(G.Mul(l0, NN.w1))

	//a1 = G.Must(G.Add(h1, NN.b1))
	l1 = G.Must(G.Sigmoid(h1))

	// Layer 2
	h2 = G.Must(G.Mul(l1, NN.w2))
	//a2, _ = G.Add(h2, NN.b2)

	// Output
	NN.pred = G.Must(G.SoftMax(h2))
	G.Read(NN.pred, &NN.predVal)

	// Backward calculation
	losses := G.Must(G.HadamardProd(NN.pred, y))
	mean := G.Must(G.Mean(losses))
	cost := G.Must(G.Neg(mean))

	var costValue G.Value = cost.Value()
	G.Read(cost, &costValue)
	log.Printf("Validation Cost: %f", costValue)

	loss := costValue.Data().(float64)

	return float32(loss), x.Shape()[0], map[string]interface{}{"test": "evaluate"}
}

// --------------------------------------------------------------------
type nn struct {
	g      *G.ExprGraph
	w1, w2 *G.Node
	b1, b2 *G.Node

	pred    *G.Node
	predVal G.Value
}

func newNN(g *G.ExprGraph) *nn {
	w1 := G.NewMatrix(g, tensor.Float64, G.WithShape(4, 4), G.WithName("w1"), G.WithInit(G.GlorotN(1.0)))
	w2 := G.NewMatrix(g, tensor.Float64, G.WithShape(4, 3), G.WithName("w2"), G.WithInit(G.GlorotN(1.0)))
	b1 := G.NewMatrix(g, tensor.Float64, G.WithShape(4, 1), G.WithName("b1"), G.WithInit(G.GlorotN(1.0)))
	b2 := G.NewMatrix(g, tensor.Float64, G.WithShape(3, 1), G.WithName("b2"), G.WithInit(G.GlorotN(1.0)))
	return &nn{
		g:  g,
		w1: w1,
		w2: w2,
		b1: b1,
		b2: b2,
	}
}

func (m *nn) learnables() G.Nodes {
	return G.Nodes{m.w1, m.w2}
}

var g = G.NewGraph()
var NN = newNN(g)

// -------------------------------------------------------------------

func main() {

	client := &CustomClient{}
	log.Println("Starting client...")
	err := c.StartClient("127.0.0.1:8080", client)
	if err != nil {
		log.Print(err)
	}

}
