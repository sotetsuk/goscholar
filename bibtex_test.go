package goscholar

import (
	"fmt"
)

func ExampleGetBibTeX() {
	url := "https://scholar.google.co.jp/scholar.bib?q=info:e6RSJHGXItcJ:scholar.google.com/&output=citation&hl=en&ct=citation"
	bibtex, _ := getBibTeX(url)
	fmt.Println(bibtex)
	// Output:
	// @inproceedings{martens2010deep, title={Deep learning via Hessian-free optimization}, author={Martens, James}, booktitle={Proceedings of the 27th International Conference on Machine Learning (ICML-10)}, pages={735--742}, year={2010}}
}
