package website_racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string) {
	firstDuration := measureResponseTime(url1)
	secondDuration := measureResponseTime(url2)

	if firstDuration < secondDuration {
		return url1
	}
	return url2
}

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}
