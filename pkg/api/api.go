package api

import (
	"encoding/json"
	"net/http"
)

// Init инициализирует HTTP-роуты API
//
// Роуты:
//
//	GET  /api/nextdate   — расчет следующей даты
//	GET  /api/task       — получить задачу (auth)
//	GET  /api/tasks      — список задач (auth)
//	POST /api/task/done  — отметить задачу выполненной (auth)
//	POST /api/signin     — аутентификация
func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", auth(taskHandler))
	http.HandleFunc("/api/tasks", auth(tasksHandler))
	http.HandleFunc("/api/task/done", auth(DoneTaskHandler))
	http.HandleFunc("/api/signin", authHandler)
}

func writeJson(w http.ResponseWriter, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(resp)
}

func writeJsonError(w http.ResponseWriter, error error) {
	data := make(map[string]string)
	data["error"] = error.Error()
	resp, err := json.Marshal((data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(resp)

}
