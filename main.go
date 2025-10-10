package main

import (
	"fmt"
	"github.com/psurge1/GoConcurrencyProblems/producer_consumer_problem"
)

func main() {
	fmt.Printf("Running Go Concurrency Problem Tests\n\n")

	producer_consumer_problem.Test()
}
