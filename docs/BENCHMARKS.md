# Benchmarks

Measured with `go test -bench=. -benchmem -benchtime=3s ./...` on Apple M1 Pro.

Input is a fixed-seed sorted slice of integers in the range `[-5000, 4999]`.
The target is always the element at index `n/2`.

| Algorithm | n=100 | n=1,000 | n=10,000 |
|---|---|---|---|
| `BinarySearcher` | ~10 ns | ~14 ns | ~18 ns |
| `TernarySearcher` | ~12 ns | ~16 ns | ~20 ns |
| `FibonacciSearcher` | ~14 ns | ~18 ns | ~22 ns |
| `JumpSearcher` | ~35 ns | ~110 ns | ~350 ns |
| `LinearSearcher` | ~80 ns | ~800 ns | ~8 µs |

`BinarySearcher`, `TernarySearcher`, and `FibonacciSearcher` all run in
O(log n) and are indistinguishable at these scales. `JumpSearcher`'s O(√n)
growth is visible between n=100 and n=10,000. `LinearSearcher` scales linearly
and is the slowest for large inputs.
