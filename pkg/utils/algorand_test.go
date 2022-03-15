package utils

import (
	"testing"

	"github.com/Dan6erbond/algogovernance/pkg/constants"
)

func TestMicroAlgoToAlgo(t *testing.T) {
	microAlgo := float64(10 / constants.MICRO_ALGO)
	algo := MicroAlgoToAlgo(microAlgo)
	if algo != 10 {
		t.Errorf("Expected 10, got %f", algo)
	}
}
