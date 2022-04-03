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

/* API Links */
const (
	ALGORAND_GOVERNANCE_API_URL                         = "https://governance.algorand.foundation/api/"
	ALGORAND_GOVERNANCE_BANNERS_PATH                    = "banners/active/"
	ALGORAND_GOVERNANCE_GOVERNOR_STATUS_PATH            = "governors/%s/status/"
	ALGORAND_GOVERNANCE_PERIODS_PATH                    = "periods/"
	ALGORAND_GOVERNANCE_ACTIVE_PERIOD_PATH              = "periods/active/"
	ALGORAND_GOVERNANCE_PERIOD_STATISTICS_PATH          = "periods/statistics/"
	ALGORAND_GOVERNANCE_PERIOD_PATH                     = "periods/%s/"
	ALGORAND_GOVERNANCE_PERIOD_GOVERNORS_PATH           = "periods/%s/governors/"
	ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_PATH            = "periods/%s/governors/%s/"
	ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_ACTIVITIES_PATH = "periods/%s/governors/%s/activities/"
	ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_STATUS_PATH     = "periods/%s/governors/%s/status/"
	ALGORAND_GOVERNANCE_TOPIC_OPTION_VOTES_PATH         = "topic-options/%s/votes/"
	ALGORAND_GOVERNANCE_TRANSACTION_PATH                = "transactions/%s/"
	ALGORAND_GOVERNANCE_VOTING_SESSION_PATH             = "voting-sessions/%s/"
)

/* GET JSON REQUESTS */

// Get makes a Get request to the given url and unmarshal the response into the given result.
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

// GetPath makes a Get request to the given path and appends any query parameters if given which are then forwarded to client.Get().
func GetPath(path string, query url.Values, result interface{}) error {
	url := ALGORAND_GOVERNANCE_API_URL + path

	if query.Encode() != "" {
		url += "?" + query.Encode()
	}

	return Get(url, result)
}

// GetBanners returns any active banners which are displayed on the governance website.
func GetBanners() (result ActiveBanners, err error) {
	err = GetPath(ALGORAND_GOVERNANCE_BANNERS_PATH, nil, &result)

	return result, err
}

// GetGovernorStatus returns a GovernorStatus object for the given address.q
func GetGovernorStatus(governor string) (result GovernorStatus, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_GOVERNOR_STATUS_PATH, governor), nil, &result)

	return result, err
}

// GetPeriods returns a paginated list of all governance periods.
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

// GetActivePeriod returns the current active governance period.
func GetActivePeriod() (result Period, err error) {
	err = GetPath(ALGORAND_GOVERNANCE_ACTIVE_PERIOD_PATH, nil, &result)

	return result, err
}

// GetPeriodStatistics returns statistics for governance such as unique governors and total rewards distributed.
// It also returns a list of past periods from which the statistics were calculated.
func GetPeriodStatistics() (result PeriodStatistics, err error) {
	err = GetPath(ALGORAND_GOVERNANCE_PERIOD_STATISTICS_PATH, nil, &result)

	return result, err
}

// GetPeriod returns a governance period by slug.
func GetPeriod(periodSlug string) (result Period, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_PATH, periodSlug), nil, &result)

	return result, err
}

// GetPeriodGovernors returns a paginated list of all governors in a governance period.
// paginator can be set to "cursor" to use cursor based pagination, which will remove the count attribute of the JSON return value.
// GetNext() and GetPrevious() can be used to navigate through the list of governors.
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

// GetPeriodGovernor returns information about a governor in a governance period.
func GetPeriodGovernor(periodSlug, governor string) (result PeriodGovernor, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_PATH, periodSlug, governor), nil, &result)

	return result, err
}

// GetPeriodGovernorStatus returns only bare data about a governor in a governance period.
// This is useful for checking if a governor is eligible to vote in a period, their committed stake, and uncompleted voting sessions.
// For a more detailed overview of the governor, use GetPeriodGovernor().
func GetPeriodGovernorStatus(periodSlug string, governor string) (result PeriodGovernorStatus, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_STATUS_PATH, periodSlug, governor), nil, &result)

	return result, err
}

// GetGovernorActivities returns a paginated list of all activities for a given governor.
// For more info, see client.GovernorActivities.
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

// GetTopicOptionVotes returns a paginated list of all the votes casted for a topic option.
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

// GetTransaction returns transaction details for the given ID.
func GetTransaction(id string) (result Transaction, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_TRANSACTION_PATH, id), nil, &result)

	return result, err
}

// GetVotingSession returns details about a voting session including topics and topic options.
func GetVotingSession(sessionSlug string) (result VotingSessionDetail, err error) {
	err = GetPath(fmt.Sprintf(ALGORAND_GOVERNANCE_VOTING_SESSION_PATH, sessionSlug), nil, &result)

	return result, err
}

/* DOWNLOADS */

// Download downloads a file from the Algorand Governance API by its path and saves it to the given file path.
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

// DownloadGovernors downloads and saves the list of governors for a governance period to the given file path.
func DownloadGovernors(periodSlug, filepath string) (err error) {
	url := ALGORAND_GOVERNANCE_API_URL + fmt.Sprintf(ALGORAND_GOVERNANCE_PERIOD_GOVERNOR_STATUS_PATH, periodSlug, "")
	return Download(url, filepath)
}
