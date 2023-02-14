package goflwr

import (
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/mat"
	"gorgonia.org/tensor"
)

type matrix struct {
	dataframe.DataFrame
}

func (m matrix) At(i, j int) float64 {
	return m.Elem(i, j).Float()
}

func (m matrix) T() mat.Matrix {
	return mat.Transpose{Matrix: m}
}

func GetXYMat() (*tensor.Dense, *tensor.Dense) {
	f, err := os.Open("c2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	df := dataframe.ReadCSV(f)

	yDF := df.Select([]string{"setosa", "versicolor", "virginica"})
	xDF := df.Drop([]string{"setosa", "versicolor", "virginica"})

	return tensor.FromMat64(mat.DenseCopyOf(&matrix{xDF})), tensor.FromMat64(mat.DenseCopyOf(&matrix{yDF}))
}
