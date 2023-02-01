package main

import "net/http"

var limitChan = make(chan struct{}, 0)

func limitMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limitChan <- struct{}{}
		next.ServeHTTP(w, r)
		<-limitChan
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello wdf"))
}

func main() {
	http.Handle("/hello", limitMiddleWare(http.HandlerFunc(hello)))

	http.ListenAndServe(":5656", nil)
}
