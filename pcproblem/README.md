### Example Output

$ go run .
Running Go Concurrency Problem Tests

Producer Consumer Problem: (16 producers) (16 consumers) (4 spots in buffer)

Produce data: {1}
Queue After Produce:
1 0 0 0


Consume data: {1}
Queue After Consume:
0 0 0 0


Produce data: {2}
Queue After Produce:
0 2 0 0


Produce data: {3}
Queue After Produce:
0 2 3 0


Produce data: {4}
Queue After Produce:
0 2 3 4


Consume data: {2}
Queue After Consume:
0 0 3 4


Produce data: {5}
Queue After Produce:
5 0 3 4


Consume data: {3}
Queue After Consume:
5 0 0 4


Consume data: {4}
Queue After Consume:
5 0 0 0


Consume data: {5}
Queue After Consume:
0 0 0 0


Produce data: {6}
Queue After Produce:
0 6 0 0


Consume data: {6}
Queue After Consume:
0 0 0 0


Produce data: {7}
Queue After Produce:
0 0 7 0


Produce data: {8}
Queue After Produce:
0 0 7 8


Consume data: {7}
Queue After Consume:
0 0 0 8


Produce data: {9}
Queue After Produce:
9 0 0 8


Consume data: {8}
Queue After Consume:
9 0 0 0


Consume data: {9}
Queue After Consume:
0 0 0 0


Produce data: {10}
Queue After Produce:
0 10 0 0


Produce data: {11}
Queue After Produce:
0 10 11 0


Produce data: {12}
Queue After Produce:
0 10 11 12


Consume data: {10}
Queue After Consume:
0 0 11 12


Consume data: {11}
Queue After Consume:
0 0 0 12


Produce data: {13}
Queue After Produce:
13 0 0 12


Produce data: {14}
Queue After Produce:
13 14 0 12


Consume data: {12}
Queue After Consume:
13 14 0 0


Produce data: {15}
Queue After Produce:
13 14 15 0


Produce data: {16}
Queue After Produce:
13 14 15 16


Consume data: {13}
Queue After Consume:
0 14 15 16


Consume data: {14}
Queue After Consume:
0 0 15 16


Consume data: {15}
Queue After Consume:
0 0 0 16


Consume data: {16}
Queue After Consume:
0 0 0 0