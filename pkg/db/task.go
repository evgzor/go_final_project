package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat,omitempty"`
}

const MaxItems = 50

func AddTask(task *Task) (int64, error) {
	var id int64
	// определите запрос
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`
	res, err := db.Exec(query, /*передайте параметры task.Date, task.Title и т.д.*/
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT %d", limit))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return processData(rows)
}

func SearchByStringTasks(search string, limit int) ([]*Task, error) {
	searchPattern := "%" + search + "%"

	query := `SELECT * FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY date LIMIT :limit`

	rows, err := db.Query(query, sql.Named("search", searchPattern), sql.Named("limit", limit))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return processData(rows)

}

func SearchByDateTasks(date string, limit int) ([]*Task, error) {

	_, err := time.Parse("20060102", date)

	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM scheduler WHERE date = :date ORDER BY date LIMIT :limit`

	rows, err := db.Query(query, sql.Named("date", date), sql.Named("limit", limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return processData(rows)
}

func processData(rows *sql.Rows) ([]*Task, error) {
	tasks := make([]*Task, 0)
	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
