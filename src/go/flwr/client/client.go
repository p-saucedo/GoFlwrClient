package goflwr

type iClient interface {
	GetParameters()
	GetProperties()
	Fit()
	Evaluate()
}
