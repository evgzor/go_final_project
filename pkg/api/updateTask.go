package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/evgzor/go_final_project/pkg/db"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

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

	err = validate(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, err)
		return
	}

	err = db.UpdateTask(&task)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, err)
		return
	}

	writeJson(w, map[string]any{})
}

func validate(task db.Task) error {
	if task.ID == "" {
		return errors.New("id is required")
	}

	if task.Title == "" {
		return errors.New("title is required")
	}

	if len(task.Repeat) > 0 {
		rules := strings.Split(task.Repeat, " ")
		mainRule := []rune(rules[0])
		symbol := mainRule[0]

		if symbol != 'y' &&
			symbol != 'd' &&
			symbol != 'w' &&
			symbol != 'm' {
			return errors.New("Wrong format data non correct charecter")
		}

		if symbol == 'y' && len(rules) > 1 {
			return errors.New("Wrong format data for year")
		}
		if len(rules) > 1 {

			_, err := strconv.Atoi(rules[1])
			if err != nil {
				return fmt.Errorf("Wrong format data for %v", symbol)
			}
		}
	}

	_, err := time.Parse("20060102", task.Date)
	if err != nil {
		return errors.New("Date is wrong format")
	}

	return nil
}
