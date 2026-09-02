package multithreadedmatrixmultiplication

// MatrixOptimal is a matrix represented as a continuous slice in memory
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
