package multithreadedmatrixmultiplication

import (
	"testing"
)

var (
	A = 2
	B = 2
	C = 2
)

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

func BenchmarkMatmul2DThreaded(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrix2D[int32](A, B)
		two := NewMatrix2D[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulThreaded(one, two)
	}
}

func BenchmarkMatmul2DThreadedRows(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrix2D[int32](A, B)
		two := NewMatrix2D[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulThreadedRows(one, two)
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

func BenchmarkMatmulOptimalThreaded(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrixOptimal[int32](A, B)
		two := NewMatrixOptimal[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulThreaded(one, two)
	}
}

func BenchmarkMatmulOptimalThreadedRows(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		one := NewMatrixOptimal[int32](A, B)
		two := NewMatrixOptimal[int32](B, C)
		RandomizeMatrix(one)
		RandomizeMatrix(two)
		b.StartTimer()

		MatmulThreadedRows(one, two)
	}
}
