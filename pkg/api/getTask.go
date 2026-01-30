package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/evgzor/go_final_project/pkg/db"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJsonError(w, errors.New("Only Get supports"))
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, errors.New("Не указан идентификатор"))
		return
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, err)
		return
	}

	task, err := db.GetTask(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJsonError(w, err)
		return
	}

	writeJson(w, task)
}
