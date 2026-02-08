package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/evgzor/go_final_project/pkg/db"
)

// GetTaskHandler godoc
// @Summary      Получить задачу
// @Description  Возвращает задачу по её ID
// @Tags         tasks
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id   query     int     true  "ID задачи"
// @Success      200  {object}  db.Task
// @Failure      400  {object}  map[string]string  "Некорректный ID"
// @Failure      401  {object}  map[string]string  "Неавторизован"
// @Failure      405  {object}  map[string]string  "Метод не поддерживается"
// @Failure      500  {object}  map[string]string  "Внутренняя ошибка"
// @Router       /api/task [get]
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJsonError(w, errors.New("Only Get supports"))
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, errors.New("id is empty"))
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
