package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/evgzor/go_final_project/pkg/db"
)

// DeleteTaskHandler godoc
// @Summary      Удалить задачу
// @Description  Удаляет задачу по ID
// @Tags         tasks
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id   query     int     true  "ID задачи"
// @Success      200  {object}  map[string]any  "Задача удалена"
// @Failure      400  {object}  map[string]string  "Некорректный ID"
// @Failure      401  {object}  map[string]string  "Неавторизован"
// @Failure      500  {object}  map[string]string  "Внутренняя ошибка"
// @Router       /api/task [delete]
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJsonError(w, errors.New("Only Delete supports"))
		return
	}
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
