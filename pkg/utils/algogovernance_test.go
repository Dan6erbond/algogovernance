package utils

import "testing"

func TestGetCursor(t *testing.T) {
	url := "https://governance.algorand.foundation/api/periods/governance-period-2/governors/?cursor=cD0yMDIxLTEyLTI0KzE2JTNBMjAlM0EwMS43NjYxMzklMkIwMCUzQTAw&paginator=cursor"
	cursor, err := GetCursor(url)
	if err != nil {
		t.Error(err)
	}
	if cursor != "cD0yMDIxLTEyLTI0KzE2JTNBMjAlM0EwMS43NjYxMzklMkIwMCUzQTAw" {
		t.Errorf("Expected cD0yMDIxLTEyLTI0KzE2JTNBMjAlM0EwMS43NjYxMzklMkIwMCUzQTAw, got %s", cursor)
	}
}
