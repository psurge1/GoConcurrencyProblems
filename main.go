package main

import (
	"fmt"

	_ "net/http/pprof"

	"github.com/psurge1/GoConcurrencyProblems/compilersteps"
	lamport "github.com/psurge1/GoConcurrencyProblems/lamports_mutual_exclusion_algorithm"
	matmul "github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication"
	"github.com/psurge1/GoConcurrencyProblems/producer_consumer_problem"
	sandbox "github.com/psurge1/GoConcurrencyProblems/sandbox"
)

var (
	pcpRun      = false
	compilerRun = true
	matmulRun   = false
	lamportRun  = false
	sandBox     = false
)

func main() {
	if pcpRun {
		fmt.Printf("Running Go Concurrency Problem Tests\n\n")
		producer_consumer_problem.Test()
	}

	if compilerRun {
		fmt.Printf("Running Token Test\n")
		compilersteps.TestToken()

		fmt.Printf("\nRunning Lexer\n")
		compilersteps.TestLexer()

		fmt.Printf("\nRunning Parser\n")
		compilersteps.TestParser()
	}

	if matmulRun {
		// run benchmark tests from inside the multithreaded_matrix_multiplication directory with go test -bench=.
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
		if multCellTh, err := matmul.MatmulCellThreaded(one, two); err == nil {
			matmul.PrintMatrix(multCellTh)
		} else {
			fmt.Println(err)
		}
		if multRowTh, err := matmul.MatmulRowThreaded(one, two); err == nil {
			matmul.PrintMatrix(multRowTh)
		} else {
			fmt.Println(err)
		}
	}

	if lamportRun {
		fmt.Printf("Running Lamport's Mutual Exclusion Algorithm Simulation")
		const NumProcesses = 50
		lamport.RunSimulation(NumProcesses)
	}

	if sandBox {
		value := sandbox.T(10)
		fmt.Printf("%d\n", value)
		value.Mutate(15)
		fmt.Printf("Value Mutated!\n")
		fmt.Printf("%d\n", value)
		fmt.Printf("address: %v", &value)

		fmt.Println("\nTesting Struct Printing")
		node := compilersteps.SIfNode{nil, nil, nil}
		fmt.Printf("%+v\n", node)
		fmt.Printf("%#v\n", node)
	}
}
