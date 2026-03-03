# 🔧 Hew CLI

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Cobra](https://img.shields.io/badge/CLI-Cobra-616ae5?style=for-the-badge)](https://github.com/spf13/cobra)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)

**Hew** is the Swiss Army knife for developers. A high-performance CLI built in Go to automate repetitive daily tasks—from UUID generation to JWT manipulation and SSH key management.

## 💡 Why "Hew"?

To **hew** means to shape something with precision — cutting exactly what needs to be cut, nothing more.

That is what this tool does: removes friction from your daily workflow with sharp, focused commands.

Short, easy to type, and impossible to forget.

## 📦 Installation

```sh
go install github.com/wallanaq/hew@latest
```

Or download a pre-built binary from the [Releases](https://github.com/wallanaq/hew/releases) page.

## 🛠 Commands

Hew is organized into intuitive subcommands. Below are the core system commands:

```
Sharp tools for developers. Hew through repetitive tasks.

Usage:
  hew [flags]
  hew [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version number of Hew

Flags:
  -h, --help      help for hew
  -v, --version   version for hew

Use "hew [command] --help" for more information about a command.
```

## 👨‍💻 Development

To start developing, you will need Go (version 1.26+).

### Getting started

Clone the repository and install dependencies:

```sh
git clone https://github.com/wallanaq/hew.git
cd hew
go mod tidy
```

Run the CLI directly without building a binary:

```sh
# Run the root command
go run ./cmd/main.go

# Run a specific subcommand
go run ./cmd/main.go [command]
```

### Makefile

This project uses a `Makefile` to automate common tasks. Here are the main commands:

- `make test`: Runs all tests.
- `make lint`: Checks the code for errors and style issues.
- `make build`: Creates a local build (snapshot) in the `dist/` directory.
- `make clean`: Cleans up build artifacts.

To see all available commands, run:

```sh
make help
```

## 🏗️ Contributors

This project exists thanks to all the people who contribute. [Contribution guide](CONTRIBUTING.md).
