package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/evgzor/go_final_project/pkg/db"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, errors.New("Id is empty"))
		return
	}

	_, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, err)
		return
	}

	err = db.DeleteTask(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJsonError(w, err)
		return
	}

	writeJson(w, map[string]any{})
}
