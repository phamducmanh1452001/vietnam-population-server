package router

import (
	"fmt"
	"net/http"
	"strings"
)

type Handle func(*ResponseWriter, *http.Request)

type Router struct {
	mux map[string]Handle
}

func NewRouter() *Router {
	return &Router{
		mux: make(map[string]Handle),
	}
}

func (router *Router) Add(path string, handle Handle) {
	router.mux[path] = handle
}

func GetHeader(url string) string {
	sl := strings.Split(url, "/")
	head := strings.Join(sl[1:], "/")
	return fmt.Sprintf("/%s", head)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head := GetHeader(r.URL.Path)

	handle, ok := router.mux[head]
	if ok {
		handleWriter := &ResponseWriter{
			size:   0,
			status: 404,
			writer: w,
		}
		handle(handleWriter, r)
	} else {
		http.NotFound(w, r)
	}
}
