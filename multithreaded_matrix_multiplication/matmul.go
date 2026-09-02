package multithreaded_matrix_multiplication

/*
For a matrix multiplication to be valid, the inner dimensions must be identical
The resultant matrix will have the dimensionality of the outer dimensions
Associativity matters
*/

import (
	"fmt"
	"sync"
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

// Matrix2D has dimensions Rows x Cols
type Matrix2D[T Number] struct {
	rows    int
	columns int
	Data    [][]T
}

func (m *Matrix2D[T]) Get(row int, col int) T {
	return m.Data[row][col]
}

func (m *Matrix2D[T]) Set(row int, col int, value T) {
	m.Data[row][col] = value
}

func (m *Matrix2D[T]) Rows() int {
	return m.rows
}

func (m *Matrix2D[T]) Columns() int {
	return m.columns
}

func (m *Matrix2D[T]) NewEmpty(rows int, cols int) Matrix[T] {
	return &Matrix2D[T]{
		rows,
		cols,
		CreateDataMatrix2D[T](rows, cols),
	}
}

func NewMatrix2D[T Number](rows int, cols int) Matrix[T] {
	return &Matrix2D[T]{
		rows,
		cols,
		CreateDataMatrix2D[T](rows, cols),
	}
}

// MatrixOptimal had dimensions Rows x Cols
type MatrixOptimal[T Number] struct {
	rows    int
	columns int
	Data    []T
}

func (m *MatrixOptimal[T]) Get(row int, col int) T {
	idx := row*m.Columns() + col
	return m.Data[idx]
}

func (m *MatrixOptimal[T]) Set(row int, col int, value T) {
	idx := row*m.Columns() + col
	m.Data[idx] = value
}

func (m *MatrixOptimal[T]) Rows() int {
	return m.rows
}

func (m *MatrixOptimal[T]) Columns() int {
	return m.columns
}

func (m *MatrixOptimal[T]) NewEmpty(rows int, cols int) Matrix[T] {
	return NewMatrixOptimal[T](rows, cols)
}

func NewMatrixOptimal[T Number](rows int, cols int) Matrix[T] {
	return &MatrixOptimal[T]{
		rows,
		cols,
		CreateDataMatrixOptimal[T](rows, cols),
	}
}

func CreateDataMatrixOptimal[T Number](rows int, cols int) []T {
	matrix := make([]T, rows*cols)
	return matrix
}

func CreateDataMatrix2D[T Number](rows int, cols int) [][]T {
	matrix := make([][]T, rows)
	for i := range matrix {
		matrix[i] = make([]T, cols)
	}
	return matrix
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

// logic: value at index i, j of result matrix is the scalar sum of multiplying the ith row of first matrix times the jth column of second matrix
func MatmulIterative[T Number](one Matrix[T], two Matrix[T]) (Matrix[T], error) {
	if one.Columns() != two.Rows() {
		return one.NewEmpty(0, 0), fmt.Errorf("invalid multiplication: dimensionality")
	}

	result := one.NewEmpty(one.Rows(), two.Columns())

	vectorHeight := one.Columns()
	for i := range result.Rows() {
		for k := range vectorHeight {
			for j := range two.Columns() {
				result.Set(i, j, result.Get(i, j)+one.Get(i, k)*two.Get(k, j))
			}
		}
	}

	return result, nil
}

func MatmulThreaded[T Number](one Matrix[T], two Matrix[T]) (Matrix[T], error) {
	if one.Columns() != two.Rows() {
		return one.NewEmpty(0, 0), fmt.Errorf("invalid multiplication: dimensionality")
	}

	result := one.NewEmpty(one.Rows(), two.Columns())

	wg := sync.WaitGroup{}
	vectorHeight := one.Columns()
	for i := range one.Rows() {
		for k := range vectorHeight {
			wg.Add(1)
			go func(row int) {
				defer wg.Done()
				for j := range two.Columns() {
					result.Set(row, j, result.Get(row, j)+one.Get(row, k)*two.Get(k, j))
				}
			}(i)
		}
	}

	wg.Wait()

	return result, nil
}

func MatmulThreadedRows[T Number](one Matrix[T], two Matrix[T]) (Matrix[T], error) {
	if one.Columns() != two.Rows() {
		return one.NewEmpty(0, 0), fmt.Errorf("invalid multiplication: dimensionality")
	}

	result := one.NewEmpty(one.Rows(), two.Columns())

	wg := sync.WaitGroup{}
	vectorHeight := one.Columns()
	for i := range one.Rows() {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			for k := range vectorHeight {
				for j := range two.Columns() {
					result.Set(row, j, result.Get(row, j)+one.Get(row, k)*two.Get(k, j))
				}
			}
		}(i)
	}

	wg.Wait()
	return result, nil
}
