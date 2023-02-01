package main

import (
	"log"
	"net/http"
	"time"
)

func boy(w http.ResponseWriter, req *http.Request) {
	time.Sleep(1)
	w.Write([]byte("啊啊啊"))
}

func timeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		begin := time.Now()
		next.ServeHTTP(writer, request)
		elapse := time.Since(begin).Nanoseconds()
		log.Printf("use time:%v", elapse)
	})

}

type ls func(http.ResponseWriter, *http.Request)

func (b ls) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	boy(w, r)
}

//################ 自己实现一个 微型框架.,可以直接 添加中间件
type middleWareChan func(http.Handler) http.Handler
type Router struct {
	middleWareChan []middleWareChan
	mux            map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		middleWareChan: make([]middleWareChan, 0),
		mux:            make(map[string]http.Handler),
	}
}

func (router *Router) Use(m middleWareChan) {
	router.middleWareChan = append(router.middleWareChan, m)
}
func (router *Router) Add(path string, handler http.Handler) {
	var mergeHandler = handler
	for i := len(router.middleWareChan) - 1; i >= 0; i-- {
		mw := router.middleWareChan[i]
		mergeHandler = mw(mergeHandler)
	}

	router.mux[path] = mergeHandler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if handler, exists := router.mux[path]; exists {
		handler.ServeHTTP(w, r)
	} else {
		http.NotFoundHandler().ServeHTTP(w, r)
		//http.NotFoundHandler().ServeHTTP(w, r)
	}
}

func main() {
	router := NewRouter()
	router.Use(timeMiddleWare)
	router.Add("/hello", http.HandlerFunc(boy))

	http.ListenAndServe("0.0.0.0:7788", router)

}
