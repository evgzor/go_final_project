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

// UpdateTaskHandler godoc
// @Summary      Обновить задачу
// @Description  Обновляет существующую задачу по ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        task  body      db.Task  true  "Обновлённые данные задачи"
// @Success      200   {object}  map[string]any  "Задача обновлена"
// @Failure      400   {object}  map[string]string  "Ошибка валидации"
// @Failure      401   {object}  map[string]string  "Неавторизован"
// @Failure      500   {object}  map[string]string  "Внутренняя ошибка"
// @Router       /api/task [put]
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJsonError(w, errors.New("Only Put supports"))
		return
	}

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
			digits := strings.Split(rules[1], ",")

			if len(digits) < 1 {
				return errors.New("Its not a digits")
			}

			for _, digit := range digits {
				_, err := strconv.Atoi(digit)
				if err != nil {
					return fmt.Errorf("Its not a digit %s", digit)
				}
			}
		}
	}

	_, err := time.Parse(defaultDateFormat, task.Date)
	if err != nil {
		return errors.New("Date is wrong format")
	}

	return nil
}
