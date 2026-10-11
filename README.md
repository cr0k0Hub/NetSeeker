# NetSeeker

```bash
  ▄        ▄  ▄                 ▄      ▄  ▄     ▄  ▄     ▄    ▄   ▄  ▄     ▄  ▄  
▄██ ▄█▄  ▄██▄▀██  ▀███  ███▀  ▄▀██▄  ▄██▄▀██  ▄██▄▀██  ▄██  ▄██ ▄██▄▀██  ▄██▄▀██ 
 █░▀ ▀░█  █░   ▀    ▀  ░ ▀   █░  ▀██  █░   ▀   █░   ▀   █░   ░█  █░   ▀   █░   ░█
 ▒▓   ▒▓  ▒▓          ▒▓     ▒▓   ▀▀  ▒▓       ▒▓       ▒▓   ▒   ▒▓       ▒▓   ▒ 
 ▀▀   ▀▀  ▀▀▀▀▀       ▀▀      ▀▀      ▀▀▀▀▀    ▀▀▀▀▀    ▀▀▀▀▀    ▀▀▀▀▀    ▀▀▀▀▀  
 ░░   ░░  ░░          ░░        ▀░▄   ░░       ░░       ░░   ░   ░░       ░░  ░  
 ▒▒   ▒▒  ▒▒          ▒▒          ▒▒  ▒▒       ▒▒       ▒▒   ▒▒  ▒▒       ▒▒  ▒▒ 
 ▓▓   ▓▓  ▓▓▄  ▓▓     ▓▓▄    ▓▓▄  ▓▓  ▓▓▄  ▓▓  ▓▓▄  ▓▓  ▓▓   ▓▓  ▓▓▄  ▓▓  ▓▓  ▓▓ 
▀██  ▀██   ▀██▄▀      █▀      ▀██▄▀    ▀██▄▀    ▀██▄▀  ▀██  ▀██   ▀██▄▀  ▀██ ▀██ 
  ▀    ▀     ▀       ▀          ▀        ▀        ▀      ▀    ▀     ▀      ▀   ▀ 
```

**NetSeeker** is a TCP port scanning tool written in Go.

The project is focused on exploring network programming, concurrent execution, TCP connection handling, context cancellation, and error classification in Go.

NetSeeker is a personal learning project developed to deepen my understanding of Go and network security.

## Features

- TCP port scanning
- Concurrent scanning using goroutines and workers
- Configurable connection timeout
- Context-aware operations
- TCP connection error classification
- Scan statistics
- Command-line configuration

> **Project status:** Work in progress. Features, architecture, and CLI behavior may change as development continues.

## Getting Started

### Prerequisites

- Go installed on your system
- Git

### Clone the repository

```bash
git clone https://github.com/cr0k0Hub/NetSeeker.git
cd NetSeeker
```

### Build

```bash
go build -o netseeker ./cmd/netseeker
```

### Run

```bash
./netseeker
```

## Usage

NetSeeker is designed to be configured through command-line arguments.

Typical configuration options include:

| Option | Description |
|--------|-------------|
| `-n` | Number of targets to scan |
| `-p` | TCP port to check |
| `-t` | Number of concurrent workers |
| `-timeout` | Connection timeout |

For the exact syntax and default values, run:

```bash
./netseeker -h
```

## Architecture

The project follows a package-oriented structure that separates the application entry point, domain types, and scanning logic.

```text
NetSeeker/
├── cmd/
│   └── netseeker/
├── internal/
│   ├── domain/
│   └── scanner/
├── go.mod
└── README.md
```

- **`cmd/netseeker`** — application entry point and command-line configuration.
- **`internal/domain`** — domain types used to represent scan targets and results.
- **`internal/scanner`** — TCP scanning logic and concurrent workers.

The architecture is evolving as the project grows.

## Learning Goals

NetSeeker is an opportunity to practice and explore:

- Go concurrency with goroutines and synchronization primitives
- Worker pool patterns
- TCP networking with the `net` package
- Context propagation and cancellation
- Timeouts and network error handling
- Atomic counters and concurrent statistics
- Separation of responsibilities between packages
- CLI design and maintainable project structure
- Git workflows and incremental development

## Inspiration

NetSeeker was inspired by **[LiquidNet](https://git.churchofmalware.org/K3ysTr0K3R/LiquidNet/src/branch/main/liquidnet.go)** by [K3ysTr0K3R](https://git.churchofmalware.org/K3ysTr0K3R).

The project served as a starting point for exploring concurrent TCP scanning in Go. NetSeeker has since evolved through its own implementation, architectural changes, and additional learning experiments.

Credit goes to the original author for the inspiration.

## Responsible Use

NetSeeker is intended for educational purposes and authorized network testing.

Only scan systems and networks that you own or have explicit permission to test. Unauthorized scanning may violate applicable laws, organizational policies, or network terms of service.

## License

No license has been specified yet.

Until a license is added, permission to use, modify, or redistribute this project should not be assumed.
