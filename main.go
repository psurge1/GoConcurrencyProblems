package main

import (
	"fmt"

	"github.com/psurge1/GoConcurrencyProblems/lexer"
	matmul "github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication"
	"github.com/psurge1/GoConcurrencyProblems/producer_consumer_problem"
)

var (
	pcpRun    = false
	lexRun    = false
	matmulRun = true
)

func main() {
	if pcpRun {
		fmt.Printf("Running Go Concurrency Problem Tests\n\n")
		producer_consumer_problem.Test()
	}

	if lexRun {
		fmt.Printf("Running Lexer\n\n")
		lexer.Test()
	}

	if matmulRun {
		fmt.Printf("Running Matmul Benchmarking\n\n")
		one := matmul.NewMatrixOptimal[int32](3, 3)
		two := matmul.NewMatrixOptimal[int32](3, 3)
		matmul.RandomizeMatrix(one)
		matmul.RandomizeMatrix(two)

		matmul.PrintMatrix(one)
		matmul.PrintMatrix(two)

		if multIt, err := matmul.MatmulIterative(one, two); err == nil {
			matmul.PrintMatrix(multIt)
		} else {
			fmt.Println(err)
		}
		if multTh, err := matmul.MatmulThreaded(one, two); err == nil {
			matmul.PrintMatrix(multTh)
		} else {
			fmt.Println(err)
		}
	}
}
