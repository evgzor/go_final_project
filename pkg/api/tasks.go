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

const maxLimitRecords = 50

// tasksHandler godoc
// @Summary      Получить список задач
// @Tags         tasks
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   db.Task
// @Failure      401  {object}  map[string]string
// @Router       /api/tasks [get]
func tasksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJsonError(w, errors.New("Only Get supports"))
		return
	}

	search := r.URL.Query().Get("search")

	switch {
	case search == "":
		tasks, err := db.Tasks(maxLimitRecords) // в параметре максимальное количество записей
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
		date := t.Format(defaultDateFormat)
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
