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

# format the code
[group('dev')]
fmt:
    golangci-lint fmt

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
