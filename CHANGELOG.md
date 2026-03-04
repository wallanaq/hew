# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.2.0] - 2026-03-03

### Added

- `version` command now checks for newer releases automatically on startup and notifies the user via stderr
- `--no-update-check` flag to skip the update check when running `hew version`
- `--debug` flag to the root command to enable verbose debug logging
- Unit tests for `cmd/root`, `cmd/version`, and `internal/version`
- `CONTRIBUTING.md` with contribution guidelines

### Changed

- Improved `README.md` with full usage documentation, development setup, and Makefile reference

## [v0.1.0] - 2026-03-03

### Added

- Initial project structure with Go module (`github.com/wallanaq/hew`)
- `version` command (`hew version`) with `--short` and `--json` flags
- `Makefile` with targets: `test`, `lint`, `build`, `release`, `clean`, `help`
- GitHub Actions workflow for automated releases via GoReleaser (triggered on `v*.*.*` tags)
- GoReleaser configuration with cross-compilation for Linux, macOS, and Windows (amd64/arm64)
- `.golangci.yaml` linter configuration
- `LICENSE` (MIT)

[Unreleased]: https://github.com/wallanaq/hew/compare/v0.2.0...HEAD
[v0.2.0]: https://github.com/wallanaq/hew/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/wallanaq/hew/releases/tag/v0.1.0
