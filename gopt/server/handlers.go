package server

import (
	"encoding/json"
	"fmt"
	"gopt/datastore"
	"net/http"
	"strconv"
	"time"
)

func getProgress(w http.ResponseWriter, r *http.Request, d datastore.Datastore) {
	key := r.Header.Get("X-key")

	if pro := d.Get(key); pro != nil {
		j, err := json.Marshal([]*datastore.Progress{pro})
		if err != nil {
			fmt.Fprint(w, "[]")
			return
		}

		fmt.Fprintf(w, "%s", string(j))
	} else {
		fmt.Fprint(w, "[]")
	}
}

func saveProgress(w http.ResponseWriter, r *http.Request, d datastore.Datastore) {
	key := r.Header.Get("X-key")

	bookIndex, _ := strconv.Atoi(r.FormValue("bookIndex"))
	progress, _ := strconv.Atoi(r.FormValue("progress"))
	//title := r.PostFormValue("title")
	deviceName := r.FormValue("deviceName")
	percentage, _ := strconv.Atoi(r.FormValue("percentage"))
	//userId := r.PostFormValue("userId")
	//accessKey := r.PostFormValue("accessKey")

	d.Put(key, &datastore.Progress{
		bookIndex,
		progress,
		percentage,
		time.Now().Format(time.RFC3339),
		deviceName,
	})
}
