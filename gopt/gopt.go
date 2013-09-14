package gopt

import (
	"net/http"
	"gopt/server"
	"gopt/datastore"
)

func init() {
	http.HandleFunc("/", server.Builder(new(datastore.Appengine)))
}
