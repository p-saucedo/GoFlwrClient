package goflwr

type Client struct {
}

func (c Client) GetParameters() []float64 {
	return make([]float64, 2)
}
