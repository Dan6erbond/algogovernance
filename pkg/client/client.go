package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	ALGORAND_GOVERNANCE_API_URL                         = "https://governance.algorand.foundation/api/"
	ALGORAND_GOVERNANCE_BANNERS_PATH                    = "banners/active/"
	ALGORAND_GOVERNANCE_GOVERNOR_STATUS_PATH            = "governors/%s/status/"
	ALGORAND_GOVERNANCE_PERIODS_PATH                    = "periods/"
	ALGORAND_GOVERNANCE_ACTIVE_PERIOD_PATH              = "periods/active/"
	ALGORAND_GOVERNANCE_PERIOD_STATISTICS_PATH          = "periods/statistics/"
	ALGORAND_GOVERNANCE_PERIOD_PATH                     = "periods/%s/"
	ALGORAND_GOVERNANCE_PERIOD_GOVERNORS_PATH           = "periods/%s/governors/"
	ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_PATH            = "periods/%s/governors/%s"
	ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_ACTIVITIES_PATH = "periods/%s/governors/%s/activities/"
	ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_STATUS_PATH     = "periods/%s/governors/%s/status/"
	ALGORAND_GOVERNANCE_TOPIC_OPTION_VOTES_PATH         = "topic-options/%s/votes/"
	ALGORAND_GOVERNANCE_TRANSACTION_PATH                = "transactions/%s"
	ALGORAND_GOVERNANCE_VOTING_SESSION_PATH             = "voting-sessions/%s"
)

/* GET JSON REQUESTS */

func Get(url string, result interface{}) error {
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

	//lint:ignore SA9003 This error should be handled in the future.
	if err != nil {
		// TODO: Handle error.
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

func GetPath(path string, query url.Values, result interface{}) error {
	url := ALGORAND_GOVERNANCE_API_URL + path

	if query.Encode() != "" {
		url += "?" + query.Encode()
	}

	return Get(url, result)
}

func GetBanners() (result ActiveBanners, err error) {
	err = GetPath(ALGORAND_GOVERNANCE_BANNERS_PATH, nil, &result)

	return result, err
}

func GetGovernorStatus(governor string) (result GovernorStatus, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_GOVERNOR_STATUS_PATH, governor), nil, &result)

	return result, err
}

func GetPeriods(limit, offset string) (result Periods, err error) {
	query := url.Values{}
	if limit != "" {
		query.Add("limit", limit)
	}
	if offset != "" {
		query.Add("offset", offset)
	}
	err = GetPath(ALGORAND_GOVERNANCE_PERIODS_PATH, query, &result)

	return result, err
}

func GetActivePeriod() (result Period, err error) {
	err = GetPath(ALGORAND_GOVERNANCE_ACTIVE_PERIOD_PATH, nil, &result)

	return result, err
}

func GetPeriodStatistics() (result PeriodStatistics, err error) {
	err = GetPath(ALGORAND_GOVERNANCE_PERIOD_STATISTICS_PATH, nil, &result)

	return result, err
}

func GetPeriod(periodSlug string) (result Period, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_PATH, periodSlug), nil, &result)

	return result, err
}

func GetPeriodGovernors(periodSlug, isEligible, address, ordering, paginator, cursor, limit, offset string) (result PeriodGovernors, err error) {
	query := url.Values{}
	if isEligible != "" {
		query.Add("is_eligible", isEligible)
	}
	if address != "" {
		query.Add("address", address)
	}
	if ordering != "" {
		query.Add("ordering", ordering)
	}
	if paginator != "" {
		query.Add("paginator", paginator)
		if cursor != "" {
			query.Add("cursor", cursor)
		}
	} else {
		if limit != "" {
			query.Add("limit", limit)
		}
		if offset != "" {
			query.Add("offset", offset)
		}
	}
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_GOVERNORS_PATH, periodSlug), query, &result)

	return result, err
}

func GetPeriodGovernor(periodSlug, governor string) (result PeriodGovernor, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_PATH, periodSlug, governor), nil, &result)

	return result, err
}

func GetPeriodGovernorStatus(periodSlug string, governor string) (result PeriodGovernorStatus, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_STATUS_PATH, periodSlug, governor), nil, &result)

	return result, err
}

func GetGovernorActivities(periodSlug, governor, limit, offset string) (result GovernorActivities, err error) {
	query := url.Values{}
	if limit != "" {
		query.Add("limit", limit)
	}
	if offset != "" {
		query.Add("offset", offset)
	}
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_ACTIVITIES_PATH, periodSlug, governor), query, &result)

	return result, err
}

func GetTopicOptionVotes(id, limit, offset string) (result TopicOptionVotes, err error) {
	query := url.Values{}
	if limit != "" {
		query.Add("limit", limit)
	}
	if offset != "" {
		query.Add("offset", offset)
	}
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_TOPIC_OPTION_VOTES_PATH, id), query, &result)

	return result, err
}

func GetTransaction(id string) (result Transaction, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_TRANSACTION_PATH, id), nil, &result)

	return result, err
}

func GetVotingSession(sessionSlug string) (result VotingSessionDetail, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_VOTING_SESSION_PATH, sessionSlug), nil, &result)

	return result, err
}

/* DOWNLOADS */

func Download(path, filepath string) (err error) {
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

	out, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func DownloadGovernors(periodSlug, filepath string) (err error) {
	url := ALGORAND_GOVERNANCE_API_URL + fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_STATUS_PATH, periodSlug, "")
	return Download(url, filepath)
}
