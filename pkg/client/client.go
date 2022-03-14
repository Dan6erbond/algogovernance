package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ALGORAND_GOVERNANCE_API_URL            = "https://governance.algorand.foundation/api/"
	ALGORAND_GOVERNANCE_PERIODS_PATH       = "periods"
	ALGORAND_GOVERNANCE_PERIOD_PATH        = "periods/%s"
	ALGORAND_GOVERNANCE_ACTIVE_PERIOD_PATH = "periods/active"
	ALGORAND_GOVERNANCE_GOVERNORS_PATH     = "periods/%s/governors/%s/status"
)

func Get(path string, result interface{}) error {
	url := ALGORAND_GOVERNANCE_API_URL + path

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(bytes))
	}

	err = json.Unmarshal(bytes, result)

	return err
}

func GetPeriods() (result Periods, err error) {
	err = Get(ALGORAND_GOVERNANCE_PERIODS_PATH, &result)

	return result, err
}

func GetPeriod(periodSlug string) (result GovernancePeriod, err error) {
	err = Get(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_PATH, periodSlug), &result)

	return result, err
}

func GetActivePeriod() (result GovernancePeriod, err error) {
	err = Get(ALGORAND_GOVERNANCE_ACTIVE_PERIOD_PATH, &result)

	return result, err
}

func GetGovernors(periodSlug string, governor string) (result Governors, err error) {
	err = Get(fmt.Sprintf(ALGORAND_GOVERNANCE_GOVERNORS_PATH, periodSlug, governor), &result)

	return result, err
}
