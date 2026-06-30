# retrievium — Claude Code instructions

retrievium is a Go library: one `Searcher` interface with a family of search algorithms behind it. Correctness and simplicity over cleverness — prefer a well-tested, focused function over a clever abstraction. If the standard library does it, use it.

## Before every commit

* Run `just ci` autonomously (lint + test + build). All must pass.
* When adding a function or package, write unit tests alongside the code in the same commit.
* The public API keeps full godoc coverage — the revive `exported` rule enforces it. Document every exported symbol you add.
* Never commit until the user explicitly confirms. Propose changes as diffs, run `just ci` autonomously, then stop and wait before `git commit`.

## Code quality

* Run `just lint` before proposing a diff. Fix all lint errors before committing.
* Prefer early returns over nesting.
* Do not add error handling or fallbacks for scenarios that cannot happen.
* Comments in `internal/` and `cmd/`: only *inside* functions, and only for non-obvious logic — never a doc comment on a declaration, a file or package header, or a `doc.go`. (`//go:build`, `//go:generate`, and `//nolint` are directives, not comments, and stay.)
* Public library code (module root or `pkg/…`, where present) keeps full godoc — a doc comment on every exported symbol and the package, enforced by the revive `exported` rule.
