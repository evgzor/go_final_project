package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/evgzor/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, errors.New("Only Get supports"))
		return
	}

	search := r.URL.Query().Get("search")

	switch {
	case search == "":
		tasks, err := db.Tasks(50) // в параметре максимальное количество записей
		if err != nil {
			writeJsonError(w, err)
			return
		}
		writeJson(w, TasksResp{
			Tasks: tasks,
		})
	case isDate(search):
		t, err := time.Parse("02.01.2006", search)
		if err != nil {
			writeJsonError(w, err)
			return
		}
		date := t.Format("20060102")
		tasks, err := db.SearchByDateTasks(date, 50)
		if err != nil {
			writeJsonError(w, err)
			return
		}
		writeJson(w, TasksResp{
			Tasks: tasks,
		})
		return
	default:
		tasks, err := db.SearchByStringTasks(search, 50)
		if err != nil {
			writeJsonError(w, err)
			return
		}
		writeJson(w, TasksResp{
			Tasks: tasks,
		})

	}
}

func isDate(date string) bool {
	_, err := time.Parse("02.01.2006", date)

	if err != nil {
		return false
	}

	return true
}
