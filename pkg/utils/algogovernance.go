package utils

import netUrl "net/url"

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
