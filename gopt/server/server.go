package server

import (
	"net/http"
	"strings"
	"gopt/datastore"
)

type handlerFunc func (http.ResponseWriter, *http.Request, datastore.Datastore)

var handlers map[string]handlerFunc

func init() {
	handlers = map[string]handlerFunc{}

	AddHandler("/progress", "GET", getProgress)
	AddHandler("/progress", "POST", saveProgress)
}

func AddHandler(path, method string, h handlerFunc) {
	handlers[path[1:]+"_"+method] = h
}

func Builder(d datastore.Datastore) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path[1:], "/")

		if len(path) > 1 {
			r.Header.Add("X-key", path[1])
		}

		if handler, ok := handlers[path[0]+"_"+r.Method]; ok {
			d.SetContext(r)
			handler(w, r, d)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
