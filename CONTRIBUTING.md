# Contributing to Hew

Thank you for your interest in contributing! This document explains how to get started, the conventions we follow, and how to submit your changes.

---

## Table of Contents

- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Development Workflow](#development-workflow)
- [Commit Conventions](#commit-conventions)
- [Pull Request Guidelines](#pull-request-guidelines)
- [Reporting Issues](#reporting-issues)

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.26+
- [golangci-lint](https://golangci-lint.run/usage/install/) (linter)
- [goreleaser](https://goreleaser.com/install/) (build & release)

### Setup

```sh
# 1. Fork the repository and clone your fork
git clone https://github.com/<your-username>/hew.git
cd hew

# 2. Install dependencies
go mod download

# 3. Install Go tooling
go install golang.org/x/tools/cmd/goimports@latest

# 4. Verify everything is working
make test
make lint
```

---

## Project Structure

```
hew/
├── cmd/              # CLI entry points (Cobra commands)
│   └── <command>/    # One package per command group
├── internal/         # Private application logic (not importable externally)
│   └── <package>/
├── .github/
│   └── workflows/    # CI/CD pipelines
├── .editorconfig
├── .golangci.yaml
├── .goreleaser.yaml
├── Makefile
└── go.mod
```

The rule of thumb: `cmd/` knows about the CLI layer (flags, output, errors). `internal/` contains pure logic with no dependency on Cobra or fmt.Print calls. Keep them separate.

---

## Development Workflow

All common tasks are available via `make`. Run `make help` to see the full list.

```sh
make test      # run all tests
make lint      # run golangci-lint
make build     # local snapshot build (output: dist/)
make clean     # remove build artifacts
```

### Adding a new command

1. Create a new package under `cmd/<command>/`
2. Implement the business logic in `internal/<package>/`
3. Register the command in `cmd/root.go`
4. Write tests for both layers
5. Document the command in `README.md`

### Writing tests

- Use the standard `testing` package together with [`testify`](https://github.com/testify/testify) for assertions
- For Cobra commands, inject output via `cmd.SetOut(buf)` and `cmd.SetErr(buf)` — never rely on `os.Stdout` in tests
- Prefer table-driven tests for multiple input cases
- Aim for 80%+ coverage on `internal/` packages

---

## Commit Conventions

This project follows [Conventional Commits](https://www.conventionalcommits.org/).

```
<type>(<scope>): <short description>
```

### Types

| Type       | When to use                                     |
| ---------- | ----------------------------------------------- |
| `feat`     | A new feature                                   |
| `fix`      | A bug fix                                       |
| `docs`     | Documentation changes only                      |
| `test`     | Adding or updating tests                        |
| `refactor` | Code change that is neither a fix nor a feature |
| `chore`    | Tooling, dependencies, CI changes               |
| `perf`     | Performance improvement                         |

### Examples

```
feat(uuid): add generate command with v4 and v7 support
fix(jwt): handle expired token error gracefully
docs: add installation section to README
chore: update golangci-lint to v1.64
test(version): add unit tests for BuildInfo.String()
```

### Rules

- Use the **imperative mood** in the description: "add", not "added" or "adds"
- Keep the subject line under **72 characters**
- Do not end the subject line with a period
- Reference issues when relevant: `feat(uuid): add inspect command (closes #12)`

---

## Pull Request Guidelines

- **One PR per feature or fix** — keep changes focused and reviewable
- **All checks must pass** before requesting review: tests, lint, and build
- **Update the README** if your change adds or modifies a command
- **Write or update tests** for any logic added in `internal/`
- Fill in the PR description with context on what changed and why

```sh
# Before opening a PR, run the full check locally
make test
make lint
make build
```

---

## Reporting Issues

- Search [existing issues](https://github.com/wallanaq/hew/issues) before opening a new one
- Use a clear title that describes the problem or request
- For bugs, include: Go version (`go version`), OS, and steps to reproduce
- For feature requests, describe the use case and expected behaviour
