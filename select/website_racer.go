package website_racer

import (
	"net/http"
)

func Racer(url1, url2 string) (winner string) {
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
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
