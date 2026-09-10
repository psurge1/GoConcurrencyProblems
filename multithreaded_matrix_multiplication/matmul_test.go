package multithreadedmatrixmultiplication

import (
	"testing"
)

var (
	A = 256
	B = 256
	C = 256
)

type Data int32

type (
	matrixFactory func(rows int, cols int) Matrix[Data]
	matmulFunc    func(one Matrix[Data], two Matrix[Data]) (Matrix[Data], error)
)

func BenchmarkMatmul(b *testing.B) {
	matrixTypes := []struct {
		name string
		new  matrixFactory
	}{
		{
			name: "Optimal",
			new:  NewMatrix2D[Data],
		},
		{
			name: "2D",
			new:  NewMatrix2D[Data],
		},
	}

	matmulAlgorithms := []struct {
		name string
		fn   matmulFunc
	}{
		{
			name: "Iterative",
			fn:   MatmulIterative[Data],
		},
		{
			name: "CellThreaded",
			fn:   MatmulCellThreaded[Data],
		},
		{
			name: "RowThreaded",
			fn:   MatmulRowThreaded[Data],
		},
		{
			name: "IterativeNoLoopInterchange",
			fn:   MatmulIterativeNoLoopInterchange[Data],
		},
		{
			name: "CellThreadedNoLoopInterchange",
			fn:   MatmulCellThreadedNoLoopInterchange[Data],
		},
		{
			name: "RowThreadedNoLoopInterchange",
			fn:   MatmulRowThreadedNoLoopInterchange[Data],
		},
	}

	for _, matrixType := range matrixTypes {
		b.Run(matrixType.name, func(b *testing.B) {
			for _, matmulAlgorithm := range matmulAlgorithms {
				b.Run(matmulAlgorithm.name, func(b *testing.B) {
					for b.Loop() {
						b.StopTimer()
						one := matrixType.new(A, B)
						two := matrixType.new(B, C)
						RandomizeMatrix(one)
						RandomizeMatrix(two)
						b.StartTimer()
						matmulAlgorithm.fn(one, two)
					}
				})
			}
		})
	}
}

/*
func BenchmarkMatmul2DIterative(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrix2D[int32](A, B)
		two := NewMatrix2D[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulIterative(one, two)
	}
}

func BenchmarkMatmul2DCellThreaded(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrix2D[int32](A, B)
		two := NewMatrix2D[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulCellThreaded(one, two)
	}
}

func BenchmarkMatmul2DRowThreaded(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrix2D[int32](A, B)
		two := NewMatrix2D[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulRowThreaded(one, two)
	}
}

func BenchmarkMatmulOptimalIterative(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrixOptimal[int32](A, B)
		two := NewMatrixOptimal[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulIterative(one, two)
	}
}

func BenchmarkMatmulOptimalCellThreaded(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrixOptimal[int32](A, B)
		two := NewMatrixOptimal[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulCellThreaded(one, two)
	}
}

func BenchmarkMatmulOptimalRowThreaded(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrixOptimal[int32](A, B)
		two := NewMatrixOptimal[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulRowThreaded(one, two)
	}
}

func BenchmarkMatmulOptimalIterativeNoLoopInterchange(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrixOptimal[int32](A, B)
		two := NewMatrixOptimal[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulIterativeNoLoopInterchange(one, two)
	}
}
*/
