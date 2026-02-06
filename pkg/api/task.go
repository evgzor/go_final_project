package api

import (
	"net/http"
)

// taskHandler godoc
// @Summary      Операции с задачей
// @Description  CRUD-операции над задачей в зависимости от HTTP-метода
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
//
// @Param        id    query     int        false  "ID задачи (GET, DELETE)"
// @Param        task  body      db.Task    false  "Данные задачи (POST, PUT)"
//
// @Success      200   {object}  db.Task                  "GET"
// @Success      200   {object}  map[string]string        "POST"
// @Success      200   {object}  map[string]any           "DELETE"
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      405   {object}  map[string]string
// @Failure      500   {object}  map[string]string
//
// @Router       /api/task [get]
// @Router       /api/task [post]
// @Router       /api/task [put]
// @Router       /api/task [delete]
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodGet:
		GetTaskHandler(w, r)
	case http.MethodPut:
		UpdateTaskHandler(w, r)
	case http.MethodDelete:
		DeleteTaskHandler(w, r)
	}

}
