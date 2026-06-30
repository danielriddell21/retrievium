# list available recipes
default:
    @just --list

# build the library
[group('build')]
build:
    go build ./...

# run the tests
[group('test')]
test:
    go test ./...

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

# full gate: lint + test + build. all must pass before committing
[group('dev')]
ci: lint test build

# run the benchmarks
[group('test')]
bench:
    go test -bench=. -benchmem ./...
