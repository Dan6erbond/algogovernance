# `algogovernance`

[![Go Reference](https://pkg.go.dev/badge/github.com/Dan6erbond/algogovernance.svg)](https://pkg.go.dev/github.com/Dan6erbond/algogovernance)

A Golang API wrapper for Algorand's Governance API and CLI tool.

## Installation

Add the Go package:

```bash
$ go get github.com/Dan6erbond/algogovernance
```

Install the CLI tool:

```bash
$ go install github.com/Dan6erbond/algogovernance
```

## Quick Start

Algorand's Governance API is a RESTful API that doesn't require keys or tokens to access, which is why all its methods are available under the `pkg/client` module.

Fetching the current period and its governors:

```go
package main

import (
	"github.com/Dan6erbond/algogovernance/pkg/client"
)

func main() {
	activePeriod, _ := client.GetActivePeriod()
  governors, _ := client.GetPeriodGovernors(activePeriod.Slug, "", "", "", "cursor", "", "", "")
  for governors.HasNext() {
    governors, _ = governors.GetNext()
  }
}
```

`algogovernance` also comes with helpers to estimate rewards for a given governor and governance period. Those are found in the `pkg/rewards` module.

Getting the estimated rewards for the current period:

```go
package main

import (
	algoRewards "github.com/Dan6erbond/algogovernance/pkg/rewards"
)

func main() {
  address := "3RYOY2LTPC6GLT3ZYE4LUFGGAEMY7GRENZQO7RFNGK2LGCV77QNASK6C6Y"
  rewards, _ := algoRewards.GetRewardsForCurrentPeriod(address)
}
```

## CLI Tool

The CLI tool is built with Cobra and uses Viper configuration for governor wallet addresses. It supports the following commands:

- `cfg`: View the current configuration.
- `currentPeriod`: Get an overview of the current period and expected governance rewards.
- `rewards`: Get the expected rewards for a given period, defaults to the current active period.

## Documentation

The documentation is available at pkg.go.dev: [algogovernance command - github.com/Dan6erbond/algogovernance - pkg.go.dev](https://pkg.go.dev/github.com/Dan6erbond/algogovernance)

## License

This package is licensed under the [MIT license](./LICENSE).
