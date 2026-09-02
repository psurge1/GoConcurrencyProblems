### Results from testing

#### I ran tests using go benchmarks on varying sizes of matrices.
#### Results show that row-based concurrency is the most optimal algorithm for large matrices, whereas single-threaded iterative multiplication is optimal for smaller matrices. This is likely due to the cost of context switching and spawning go routines, which scales really well, but doesn't provide great benefits at smaller dimensions.

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
