# list available recipes
default:
    @just --list

# run the tests
[group('test')]
test:
    go test ./...

# run the benchmarks
[group('test')]
bench:
    go test -bench=. -benchmem ./...

# run the linter
[group('dev')]
lint:
    golangci-lint run

# vet the code
[group('dev')]
vet:
    go vet ./...

# format the code
[group('dev')]
fmt:
    gofmt -w .

# tidy module dependencies
[group('dev')]
tidy:
    go mod tidy
