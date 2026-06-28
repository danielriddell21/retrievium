# list available recipes
default:
    @just --list

# run the tests
test:
    go test ./...

# run the benchmarks
bench:
    go test -bench=. -benchmem ./...

# run the linter
lint:
    golangci-lint run

# vet the code
vet:
    go vet ./...

# format the code
fmt:
    gofmt -w .

# tidy module dependencies
tidy:
    go mod tidy
