### Results from testing

#### I ran tests using go benchmarks on varying sizes of matrices.
#### Results show that row-based concurrency is the most optimal algorithm for large matrices, whereas single-threaded iterative multiplication is optimal for smaller matrices. This is likely due to the cost of context switching and spawning go routines, which scales really well, but doesn't provide great benefits at smaller dimensions.
#### Furthermore, memory access plays a major role in performance. While loop interchange can enable memory-access locality (such that we consecutively reference adjacent blocks of memory), it leads to 3n reads n writes per cell (where n == width(A) == height(B) for the multiplication AxN). By contrast, non loop interchange lowers memory access-locality, but only 2b reads and only 1 write per cell.

MatmulIterative refers to single-threaded iterative multiplication
MatmulThreaded spawns one goroutine per location in the final matrix, and each goroutine computes that value
MatmulThreadedRows spawns one goroutine per row in the final matrix, and computes the row vector

2D means that the underlying memory structure of the matrix was a slice of slices.
Optimal means that the underlying memory structure of the matrix was one consecutive slice.

### On 512 x 512 multiplied by 512 x 512 matrix
go test -bench=. ./multithreaded_matrix_multiplication  
goos: darwin  
goarch: arm64  
pkg: github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication  
cpu: Apple M1 Pro  
BenchmarkMatmul2DIterative-10            	       2	 841521084 ns/op  
BenchmarkMatmul2DThreaded-10             	       5	 212127334 ns/op  
BenchmarkMatmul2DThreadedRows-10         	      10	 109624783 ns/op  
BenchmarkMatmulOptimalIterative-10       	       2	 761528396 ns/op  
BenchmarkMatmulOptimalThreaded-10        	       5	 210021475 ns/op  
BenchmarkMatmulOptimalThreadedRows-10    	      12	  97932007 ns/op  
PASS  
ok  	github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication	8.117s  

### On 2048 x 512 multiplied by 512 x 2048 matrix
go test -bench=. ./multithreaded_matrix_multiplication  
goos: darwin  
goarch: arm64  
pkg: github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication  
cpu: Apple M1 Pro  
BenchmarkMatmul2DIterative-10            	       1	13441732375 ns/op  
BenchmarkMatmul2DThreaded-10             	       1	2453246292 ns/op  
BenchmarkMatmul2DThreadedRows-10         	       1	1701694292 ns/op  
BenchmarkMatmulOptimalIterative-10       	       1	12070550375 ns/op  
BenchmarkMatmulOptimalThreaded-10        	       1	2405682375 ns/op  
BenchmarkMatmulOptimalThreadedRows-10    	       1	1563022792 ns/op  
PASS  
ok  	github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication	34.161  

### On 5 x 5 multiplied by 5 x 5 matrix
go test -bench=. ./multithreaded_matrix_multiplication  
goos: darwin  
goarch: arm64  
pkg: github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication  
cpu: Apple M1 Pro  
BenchmarkMatmul2DIterative-10            	  842422	      1363 ns/op  
BenchmarkMatmul2DThreaded-10             	  163816	      7655 ns/op  
BenchmarkMatmul2DThreadedRows-10         	  352102	      3365 ns/op  
BenchmarkMatmulOptimalIterative-10       	 1000000	      1123 ns/op  
BenchmarkMatmulOptimalThreaded-10        	  159544	      7194 ns/op  
BenchmarkMatmulOptimalThreadedRows-10    	  397598	      3075 ns/op  
PASS  
ok  	github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication	131.591s  

### Loop Interchange Enabled (no memory optimizations)
### 500 x 500 x 500
goos: darwin
goarch: arm64
pkg: github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication
cpu: Apple M1 Pro
BenchmarkMatmul/Optimal/Iterative-10         	       8	 126309208 ns/op
BenchmarkMatmul/Optimal/CellThreaded-10      	      30	  40441386 ns/op
BenchmarkMatmul/Optimal/RowThreaded-10       	      74	  18236456 ns/op
BenchmarkMatmul/Optimal/IterativeNoLoopInterchange-10         	       8	 136306552 ns/op
BenchmarkMatmul/Optimal/CellThreadedNoLoopInterchange-10      	      32	  36227378 ns/op
BenchmarkMatmul/Optimal/RowThreadedNoLoopInterchange-10       	      67	  18473001 ns/op
BenchmarkMatmul/2D/Iterative-10                               	       8	 125596849 ns/op
BenchmarkMatmul/2D/CellThreaded-10                            	      32	  41713466 ns/op
BenchmarkMatmul/2D/RowThreaded-10                             	      66	  17710883 ns/op
BenchmarkMatmul/2D/IterativeNoLoopInterchange-10              	       8	 136496886 ns/op
BenchmarkMatmul/2D/CellThreadedNoLoopInterchange-10           	      32	  36532501 ns/op
BenchmarkMatmul/2D/RowThreadedNoLoopInterchange-10            	      58	  18411279 ns/op

This implementation shows that loop interchange generally increases speed (except for the case of cell-based multithreading), likely due to memory-access locality

### Loop Interchange Enabled (memory optimized for Non-Loop Interchange)
goos: darwin
goarch: arm64
pkg: github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication
cpu: Apple M1 Pro
BenchmarkMatmul/Optimal/Iterative-10         	       8	 125904932 ns/op
BenchmarkMatmul/Optimal/CellThreaded-10      	      28	  38022156 ns/op
BenchmarkMatmul/Optimal/RowThreaded-10       	      79	  16665928 ns/op
BenchmarkMatmul/Optimal/IterativeNoLoopInterchange-10         	      14	  77977631 ns/op
BenchmarkMatmul/Optimal/CellThreadedNoLoopInterchange-10      	      38	  30520800 ns/op
BenchmarkMatmul/Optimal/RowThreadedNoLoopInterchange-10       	      96	  10607233 ns/op
BenchmarkMatmul/2D/Iterative-10                               	       8	 126540990 ns/op
BenchmarkMatmul/2D/CellThreaded-10                            	      30	  38058742 ns/op
BenchmarkMatmul/2D/RowThreaded-10                             	      81	  16456377 ns/op
BenchmarkMatmul/2D/IterativeNoLoopInterchange-10              	      15	  77652122 ns/op
BenchmarkMatmul/2D/CellThreadedNoLoopInterchange-10           	      34	  30639819 ns/op
BenchmarkMatmul/2D/RowThreadedNoLoopInterchange-10            	     100	  11038443 ns/op
PASS
ok  	github.com/psurge1/GoConcurrencyProblems/multithreaded_matrix_multiplication	14.555s

Without loop interchange, we can optimize out reads and writes by caching intermediate cell values, instead of accessing them from the matrix itself. 
In this case, not implementing loop interchange and caching cell values yielded much greater performance benefits.
