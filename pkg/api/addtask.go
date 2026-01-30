package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/evgzor/go_final_project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	var buf bytes.Buffer
	// читаем тело запроса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// десериализуем JSON в db
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, err)
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, errors.New("title is required"))
		return
	}

	if err := checkDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, err)
		return
	}

	idTask, err := db.AddTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, err)
		return
	}
	idTaskStr := strconv.Itoa(int(idTask))
	mapStr := make(map[string]string)
	mapStr["id"] = idTaskStr

	writeJson(w, mapStr)
}

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == now.Format("20060102") {
		return nil
	}

	if task.Date == "" {
		task.Date = now.Format("20060102")
		return nil
	}

	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("invalid date format")
	}

	if len(task.Repeat) > 0 {
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		if afterNow(now, t) {
			task.Date = next
		}
	} else {
		if afterNow(now, t) {
			task.Date = now.Format("20060102")
		}
	}

	return nil
}
