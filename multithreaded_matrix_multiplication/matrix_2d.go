package multithreadedmatrixmultiplication

// Matrix2D is a matrix represented as a slice of slices in memory (suboptimal)
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
	return NewMatrix2D[T](rows, cols)
}

func NewMatrix2D[T Number](rows int, cols int) Matrix[T] {
	return &Matrix2D[T]{
		rows,
		cols,
		CreateDataMatrix2D[T](rows, cols),
	}
}

func CreateDataMatrix2D[T Number](rows int, cols int) [][]T {
	matrix := make([][]T, rows)
	for i := range matrix {
		matrix[i] = make([]T, cols)
	}
	return matrix
}
