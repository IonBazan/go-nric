# Singapore NRIC/FIN Number Validator

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/ionbazan/go-nric/ci.yml?branch=main)](https://github.com/ionbazan/go-nric/actions)
[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.20-00ADD8)](https://golang.org/)
[![Codecov](https://img.shields.io/codecov/c/gh/ionbazan/go-nric)](https://codecov.io/gh/ionbazan/go-nric)
[![Go Report Card](https://goreportcard.com/badge/github.com/ionbazan/go-nric?style=flat-square)](https://goreportcard.com/report/github.com/ionbazan/go-nric)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/ionbazan/go-nric)](https://pkg.go.dev/mod/github.com/ionbazan/go-nric)
[![License](https://img.shields.io/github/license/ionbazan/go-nric)](https://github.com/ionbazan/go-nric/blob/main/LICENSE)

**Supports new 2022 M-series FIN numbers!**

This package provides a simple and efficient library for validating Singapore NRIC and FIN numbers in Go. 
It includes support for checksum validation and format checking, with clear error handling for invalid inputs.

## Installation

To use this library in your Go project, run:

```bash
go get github.com/ionbazan/go-nric
```

## Usage

The library provides a `nric.NRIC` type that validates NRIC/FIN numbers upon creation. 
It supports checking if a number is a foreigner ID, from 2000 or later, or part of the M-series (post-2022 FINs).

Here’s an example of how to use it:

```go
package main

import (
	"fmt"
	"github.com/ionbazan/go-nric"
)

func main() {
	// Validate a valid NRIC
	validNRIC, err := nric.NewNRIC("S6083480F")
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return
	}
	fmt.Printf("Valid NRIC: %s\n", validNRIC.String())
	fmt.Printf("Is Foreigner: %v\n", validNRIC.IsForeigner())
	fmt.Printf("Is 2000+: %v\n", validNRIC.Is2000())
	fmt.Printf("Is Series M: %v\n", validNRIC.IsSeriesM())

	_, err := nric.NewNRIC("S1234567E")
	if err != nil {
		switch {
		case errors.Is(err, nric.ErrInvalidFormat):
			fmt.Printf("Invalid format: %v\n", err)
		case errors.Is(err, nric.ErrInvalidChecksum):
			fmt.Printf("Invalid checksum: %v\n", err)
		default:
			fmt.Printf("Unexpected error: %v\n", err)
		}
		return
	}
}
```

### Example Output
```
Valid NRIC: S6083480F
Is Foreigner: false
Is 2000+: false
Is Series M: false
Invalid checksum: invalid NRIC checksum: S1234567E
```

## Features

- Validates NRIC and FIN numbers against format and checksum rules.
- Supports pre-2000 (S/F), post-2000 (T/G), and post-2022 (M) series.
- Provides methods to check:
    - `IsForeigner()`: Whether the ID is a FIN (foreigner).
    - `Is2000()`: Whether the ID is from 2000 or later.
    - `IsSeriesM()`: Whether the ID is from the M-series (post-2022 FINs).
- Returns specific errors (`ErrInvalidFormat`, `ErrInvalidChecksum`) for invalid inputs.

## Requirements

- Go 1.20 or later.

## Development

### Running Tests
To run the tests locally:

```bash
go test -v ./...
```

### Checking Coverage
To generate a coverage report:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

Open `coverage.html` in a browser to view the coverage details.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue on [GitHub](https://github.com/ionbazan/go-nric).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.