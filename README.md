# retrievium

*n.* the practice of getting back what you put in. Also: making your slices give up their secrets.

One interface. No haystack required.

[![Go Reference](https://pkg.go.dev/badge/github.com/danielriddell21/retrievium.svg)](https://pkg.go.dev/github.com/danielriddell21/retrievium)
[![CI](https://github.com/danielriddell21/retrievium/actions/workflows/ci.yaml/badge.svg)](https://github.com/danielriddell21/retrievium/actions/workflows/ci.yaml)
[![Go 1.26](https://img.shields.io/badge/go-1.26-blue)](https://go.dev)
[![MIT License](https://img.shields.io/badge/licence-MIT-green)](LICENSE)

## Installation

```sh
go get github.com/danielriddell21/retrievium@latest
```

## Quick start

```go
import "github.com/danielriddell21/retrievium"

s := retrievium.BinarySearcher{}
idx := s.Search([]int{1, 3, 5, 7, 9}, 7)
// idx == 3
fmt.Println(s.Name()) // "Binary Search"
```

Every searcher satisfies the `Searcher` interface:

```go
type Searcher interface {
    Search(haystack []int, target int) int
    Name() string
}
```

`Search` returns the **index** of the target, or **-1** if not present.
`BinarySearcher`, `TernarySearcher`, `FibonacciSearcher`, and `JumpSearcher` require the input to be sorted in ascending order. `LinearSearcher` works on any input.

## Algorithms

| Searcher | Time | Space | Requires sorted input |
|---|---|---|---|
| `LinearSearcher` | O(n) | O(1) | No |
| `BinarySearcher` | O(log n) | O(1) | Yes |
| `TernarySearcher` | O(log₃ n) | O(1) | Yes |
| `FibonacciSearcher` | O(log n) | O(1) | Yes |
| `JumpSearcher` | O(√n) | O(1) | Yes |

## Examples

Runnable examples for every algorithm are in the [`examples/`](examples/EXAMPLES.md) directory.

## Documentation
- [Benchmark results](docs/benchmarks.md)
- [Complexity chart](docs/complexity.md)
