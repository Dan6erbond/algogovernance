package utils

import netUrl "net/url"

// GetCursor returns a cursor from a URL's query string if any is found to work with Algorand's Governance API pagination.
func GetCursor(url string) (cursor string, err error) {
	u, err := netUrl.Parse(url)

	if err != nil {
		return "", err
	}

	q := u.Query()

	if q.Has("cursor") {
		return q.Get("cursor"), nil
	}

	return "", nil
}
