package multithreadedmatrixmultiplication

import (
	"testing"
)

var (
	A = 256
	B = 256
	C = 256
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
