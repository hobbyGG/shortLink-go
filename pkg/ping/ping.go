package ping

import (
	"net/http"
	"time"
)

func New(url string) bool {
	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}
