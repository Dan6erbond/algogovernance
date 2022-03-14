package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ALGORAND_GOVERNANCE_API_URL            = "https://governance.algorand.foundation/api/"
	ALGORAND_GOVERNANCE_PERIODS_PATH       = "periods"
	ALGORAND_GOVERNANCE_PERIOD_PATH        = "periods/%s"
	ALGORAND_GOVERNANCE_ACTIVE_PERIOD_PATH = "periods/active"
	ALGORAND_GOVERNANCE_GOVERNORS_PATH     = "periods/%s/governors/%s/status"
)

func Get(path string, query url.Values, result interface{}) error {
	url := ALGORAND_GOVERNANCE_API_URL + path

	if query.Encode() != "" {
		url += "?" + query.Encode()
	}

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

	errorResp := ErrorResponse{}
	err = json.Unmarshal(bytes, &errorResp)

	if err != nil {
		// If we can't unmarshal the error response, we can't determine if it's an error or not.
	}

	if errorResp.Type != "" {
		for _, value := range errorResp.Detail {
			return fmt.Errorf("%s: %s", errorResp.Type, value[0])
		}
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(bytes))
	}

	err = json.Unmarshal(bytes, result)

	return err
}

func GetPeriods(limit, offset string) (result Periods, err error) {
	query := url.Values{}
	if limit != "" {
		query.Add("limit", limit)
	}
	if offset != "" {
		query.Add("offset", offset)
	}
	err = Get(ALGORAND_GOVERNANCE_PERIODS_PATH, query, &result)

	return result, err
}

func GetPeriod(periodSlug string) (result GovernancePeriod, err error) {
	err = Get(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_PATH, periodSlug), nil, &result)

	return result, err
}

func GetActivePeriod() (result GovernancePeriod, err error) {
	err = Get(ALGORAND_GOVERNANCE_ACTIVE_PERIOD_PATH, nil, &result)

	return result, err
}

func GetGovernors(periodSlug string, governor string) (result Governors, err error) {
	err = Get(fmt.Sprintf(ALGORAND_GOVERNANCE_GOVERNORS_PATH, periodSlug, governor), nil, &result)

	return result, err
}
