package utils

import "github.com/Dan6erbond/algogovernance/pkg/constants"

// MicroAlgoToAlgo converts an amount of micro ALGO to the ALGO unit.
func MicroAlgoToAlgo(microAlgo float64) float64 {
	return microAlgo * constants.MICRO_ALGO
}
