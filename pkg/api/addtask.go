package api

import (
	"fmt"
	"net/http"

	"github.com/evgzor/go_final_project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func checkDate(task *db.Task) error {
	return fmt.Errorf("")
}
