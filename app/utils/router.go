package utils

import (
	"fmt"
	"net/http"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request)

type Router struct {
	mux map[string]Handle
}

func newRouter() *Router {
	return &Router{
		mux: make(map[string]Handle),
	}
}

func (router *Router) Add(path string, handle Handle) {
	router.mux[path] = handle
}

func GetHeader(url string) string {
	sl := strings.Split(url, "/")
	return fmt.Sprintf("/%s", sl[1])
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head := GetHeader(r.URL.Path)
	handle, ok := router.mux[head]
	if ok {
		handle(w, r)
	} else {
		http.NotFound(w, r)
	}
}
