# Contributing

Pull requests are welcome.

## Prerequisites

Go 1.25 or later.

## Running tests

```sh
go test ./...
```

## Running benchmarks

```sh
go test -bench=. -benchmem ./...
```

## Adding a new searcher

1. Create `<name>.go` implementing the `Searcher` interface:

```go
type MySearcher struct{}

func (MySearcher) Name() string { return "My Search" }

func (MySearcher) Search(haystack []int, target int) int {
    // ... search haystack for target ...
    // return index if found, -1 if not
    return -1
}
```

2. Create `<name>_test.go` in package `retrievium_test`. Use the shared helpers from `searcher_test.go` (`sortedSlice`). Include at minimum: empty input, single element found, single element not found, first element, last element, middle element, not found, and a random-slice test. Add a benchmark and an `Example` function.

3. Add an example in `examples/<scenario>/main.go`. Keep it runnable with `go run`. See existing examples for tone and structure.

4. Add a row to the algorithms table in `README.md` and a row in `examples/EXAMPLES.md`.
