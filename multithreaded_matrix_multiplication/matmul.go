// Package multithreadedmatrixmultiplication is a package exploring multithreaded implementations of matrix multiplication// multithreaded_matrix_multiplication is a package exploring multithreaded implementations of matrix multiplication// multithreaded_matrix_multiplication is a package exploring multithreaded implementations of matrix multiplication// multithreaded_matrix_multiplication is a package exploring multithreaded implementations of matrix multiplication// multithreaded_matrix_multiplication is a package exploring multithreaded implementations of matrix multiplication// multithreaded_matrix_multiplication is a package exploring multithreaded implementations of matrix multiplication// multithreaded_matrix_multiplication is a package exploring multithreaded implementations of matrix multiplication// multithreaded_matrix_multiplication is a package exploring multithreaded implementations of matrix multiplication
package multithreadedmatrixmultiplication

/*
For a matrix multiplication to be valid, the inner dimensions must be identical
The resultant matrix will have the dimensionality of the outer dimensions
Associativity matters

logic: value at index i, j of result matrix is the scalar sum of multiplying the ith row of first matrix times the jth column of second matrix
*/

import (
	"fmt"
	"sync"
)

func MatmulIterativeNoLoopInterchange[T Number](one Matrix[T], two Matrix[T]) (Matrix[T], error) {
	if one.Columns() != two.Rows() {
		return one.NewEmpty(0, 0), fmt.Errorf("invalid multiplication: dimensionality")
	}

	result := one.NewEmpty(one.Rows(), two.Columns())

	vectorHeight := one.Columns()
	for i := range result.Rows() {
		for j := range result.Columns() {
			res := T(0)
			for k := range vectorHeight {
				res += one.Get(i, k) * two.Get(k, j)
			}
			result.Set(i, j, res)
		}
	}

	return result, nil
}

func MatmulIterative[T Number](one Matrix[T], two Matrix[T]) (Matrix[T], error) {
	if one.Columns() != two.Rows() {
		return one.NewEmpty(0, 0), fmt.Errorf("invalid multiplication: dimensionality")
	}

	result := one.NewEmpty(one.Rows(), two.Columns())

	vectorHeight := one.Columns()
	for i := range result.Rows() {
		for k := range vectorHeight {
			for j := range result.Columns() {
				result.Set(i, j, result.Get(i, j)+one.Get(i, k)*two.Get(k, j))
			}
		}
	}

	return result, nil
}

func MatmulCellThreadedNoLoopInterchange[T Number](one Matrix[T], two Matrix[T]) (Matrix[T], error) {
	if one.Columns() != two.Rows() {
		return one.NewEmpty(0, 0), fmt.Errorf("invalid multiplication: dimensionality")
	}

	result := one.NewEmpty(one.Rows(), two.Columns())

	wg := sync.WaitGroup{}
	vectorHeight := one.Columns()
	for i := range result.Rows() {
		for j := range result.Columns() {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				res := T(0)
				for k := range vectorHeight {
					res += one.Get(i, k) * two.Get(k, j)
				}
				result.Set(i, j, res)
			}(i)
		}
	}

	wg.Wait()

	return result, nil
}

func MatmulCellThreaded[T Number](one Matrix[T], two Matrix[T]) (Matrix[T], error) {
	if one.Columns() != two.Rows() {
		return one.NewEmpty(0, 0), fmt.Errorf("invalid multiplication: dimensionality")
	}

	result := one.NewEmpty(one.Rows(), two.Columns())

	wg := sync.WaitGroup{}
	vectorHeight := one.Columns()
	for i := range result.Rows() {
		for k := range vectorHeight {
			wg.Add(1)
			go func(i int, k int) {
				defer wg.Done()
				for j := range result.Columns() {
					result.Set(i, j, result.Get(i, j)+one.Get(i, k)*two.Get(k, j))
				}
			}(i, k)
		}
	}

	wg.Wait()

	return result, nil
}

func MatmulRowThreadedNoLoopInterchange[T Number](one Matrix[T], two Matrix[T]) (Matrix[T], error) {
	if one.Columns() != two.Rows() {
		return one.NewEmpty(0, 0), fmt.Errorf("invalid multiplication: dimensionality")
	}

	result := one.NewEmpty(one.Rows(), two.Columns())

	wg := sync.WaitGroup{}
	vectorHeight := one.Columns()
	for i := range result.Rows() {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := range result.Columns() {
				res := T(0)
				for k := range vectorHeight {
					res += one.Get(i, k) * two.Get(k, j)
				}
				result.Set(i, j, res)
			}
		}(i)
	}

	wg.Wait()
	return result, nil
}

func MatmulRowThreaded[T Number](one Matrix[T], two Matrix[T]) (Matrix[T], error) {
	if one.Columns() != two.Rows() {
		return one.NewEmpty(0, 0), fmt.Errorf("invalid multiplication: dimensionality")
	}

	result := one.NewEmpty(one.Rows(), two.Columns())

	wg := sync.WaitGroup{}
	vectorHeight := one.Columns()
	for i := range one.Rows() {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for k := range vectorHeight {
				for j := range two.Columns() {
					result.Set(i, j, result.Get(i, j)+one.Get(i, k)*two.Get(k, j))
				}
			}
		}(i)
	}

	wg.Wait()
	return result, nil
}
