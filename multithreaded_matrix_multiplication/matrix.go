package multithreadedmatrixmultiplication

import (
	"fmt"
)

type Number interface {
	~int8 | ~int16 | ~int32 | ~int64
}

type Matrix[T Number] interface {
	Get(int, int) T
	Set(int, int, T)
	Rows() int
	Columns() int
	NewEmpty(int, int) Matrix[T]
}

func RandomizeMatrix[T Number](matrix Matrix[T]) {
	for i := range matrix.Rows() {
		for j := range matrix.Columns() {
			matrix.Set(i, j, T(RandInRange(-5, 5)))
		}
	}
}

func PrintMatrix[T Number](matrix Matrix[T]) {
	fmt.Printf("Matrix %d x %d\n", matrix.Rows(), matrix.Columns())
	for i := range matrix.Rows() {
		for j := range matrix.Columns() {
			fmt.Printf("%d ", matrix.Get(i, j))
		}
		fmt.Println("")
	}
	fmt.Println("")
}
