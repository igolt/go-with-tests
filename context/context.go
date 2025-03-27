package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		dataChan := make(chan string, 1)

		go func() {
			dataChan <- store.Fetch()
		}()

		select {
		case data := <-dataChan:
			fmt.Fprint(w, data)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
