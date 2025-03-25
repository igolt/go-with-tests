package website_racer

import (
	"fmt"
	"net/http"
	"time"
)

const racerDefaultTimeout = 10 * time.Second

func Racer(url1, url2 string) (winner string, error error) {
	return ConfigurableRacer(url1, url2, racerDefaultTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
