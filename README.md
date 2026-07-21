# retrievium

*n.* the practice of getting back what you put in. Also: making your slices give up their secrets.

One interface. No haystack required.

[![Go Reference](https://pkg.go.dev/badge/github.com/danielriddell21/retrievium.svg)](https://pkg.go.dev/github.com/danielriddell21/retrievium)
[![CI](https://github.com/danielriddell21/retrievium/actions/workflows/ci.yaml/badge.svg)](https://github.com/danielriddell21/retrievium/actions/workflows/ci.yaml)
[![Go 1.26](https://img.shields.io/badge/go-1.26-blue)](https://go.dev)
[![MIT License](https://img.shields.io/badge/licence-MIT-green)](LICENSE)

## Install
```sh
go get github.com/danielriddell21/retrievium/v2@latest
```

## Quick start

```go
import "github.com/danielriddell21/retrievium/v2"

s := retrievium.BinarySearcher[int]{}
idx, ok := s.Search([]int{1, 3, 5, 7, 9}, 7)
// idx == 3, ok == true
fmt.Println(s.Name()) // "Binary Search"

// Searchers are generic over any cmp.Ordered type:
i, found := retrievium.BinarySearcher[string]{}.Search([]string{"apple", "kiwi", "pear"}, "kiwi")
// i == 1, found == true
```

Every searcher satisfies the `Searcher` interface, generic over the element type:

```go
type Searcher[E cmp.Ordered] interface {
    Search(haystack []E, target E) (int, bool)
    Name() string
}
```

`Search` returns the **index** of the target and a boolean reporting whether it
was found (mirroring `slices.BinarySearch`). `BinarySearcher`, `TernarySearcher`,
`FibonacciSearcher`, and `JumpSearcher` require the input to be sorted in
ascending order. `LinearSearcher` works on any input.

## Why not `slices.BinarySearch`?

For production code, reach for the standard library — `slices.BinarySearch` and
`slices.Index` are the right tool. retrievium exists for teaching and
exploration: a single interface that lets you swap in and compare a family of
search algorithms at runtime.

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
