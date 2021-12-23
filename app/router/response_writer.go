package router

import (
	"log"
	"net/http"
)

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

type ResponseWriter struct {
	writer http.ResponseWriter
	size   int
	status int
}

func (w *ResponseWriter) Writer() http.ResponseWriter {
	return w.writer
}

func (w *ResponseWriter) WriteHeader(code int) {
	if code > 0 && w.status != code {
		if w.Written() {
			log.Printf("[WARNING] Headers were already written. Wanted to override status code %d with %d", w.status, code)
		}
		w.status = code
	}
}

func (w *ResponseWriter) WriteHeaderNow() {
	if !w.Written() {
		w.size = 0
		w.writer.WriteHeader(w.status)
	}
}

func (w *ResponseWriter) Write(data []byte) (n int, err error) {
	w.WriteHeaderNow()
	n, err = w.writer.Write(data)
	w.size += n
	return
}

func (w *ResponseWriter) Status() int {
	return w.status
}

func (w *ResponseWriter) Size() int {
	return w.size
}

func (w *ResponseWriter) Written() bool {
	return w.size != noWritten
}
