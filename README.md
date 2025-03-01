# Scanner in Golang

This is a simple website subdomains and port scanner built using Go.

## Features

- Scans for open ports on a given host.
- Scans for active subdomains on a given domain.
- Concurrent scanning for faster results.
- Logs execution time and results.
- Saves scan results to a file.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/gungindi/go_port_scanner.git
    cd go_port_scanner
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Build the project:
    ```sh
    go build -o port-scanner cmd/port-scanner/main.go
    ```

2. Run the scanner:
    ```sh
    ./port-scanner
    ```

### Command-Line Arguments

The scanner retrieves input from a configuration file or command-line arguments. Ensure you have the necessary input data for subdomains, ports, and host.

| Argument       | Description                                                    | Default                        |
|----------------|----------------------------------------------------------------|--------------------------------|
| `--host`       | Target host or domain to scan.                                 | N/A (Required Argument)        |
| `--ports`      | Comma-separated list of ports to scan (e.g., `80,22,3306`).    | `/internal/data/ports.txt`     |
| `--subdomains` | Comma-separated list of subdomains to scan (e.g., `www,api,blog`). | `/internal/data/subdomains.txt`|

## Example

```sh
./port-scanner --host example.com --ports 80,443 --subdomains www,api,blog