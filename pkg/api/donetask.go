package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/evgzor/go_final_project/pkg/db"
)

// tasksHandler godoc
// @Summary      Получить список задач
// @Tags         tasks
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   db.Task
// @Failure      401  {object}  map[string]string
// @Router       /api/tasks [get]
func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, errors.New("Id should not empty"))
		return
	}

	_, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, errors.New("Id is not a number"))
		return
	}

	task, err := db.GetTask(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJsonError(w, err)
		return
	}

	if task.Repeat == "" {
		err := db.DeleteTask(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeJsonError(w, err)
			return
		}
		writeJson(w, map[string]any{})
		return
	}

	task.Date, err = NextDate(time.Now(), task.Date, task.Repeat)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJsonError(w, err)
		return
	}

	err = db.UpdateDate(task.Date, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJsonError(w, err)
		return
	}

	writeJson(w, map[string]any{})
}
