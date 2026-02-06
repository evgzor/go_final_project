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

// GetTask возвращает задачу по её ID.
//
// Возвращает ошибку, если задача не найдена или произошла ошибка БД.
func GetTask(id string) (*Task, error) {
	row := db.QueryRow(`SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id`, sql.Named("id", id))

	task := Task{}

	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTask обновляет данные задачи по её ID.
//
// Возвращает ошибку, если задача с таким ID не существует.
func UpdateTask(task *Task) error {
	// параметры пропущены, не забудьте указать WHERE
	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id`
	res, err := db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID),
	)

	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

// UpdateDate обновляет только дату выполнения задачи.
func UpdateDate(next string, id string) error {
	query := `UPDATE scheduler SET date = :date WHERE id = :id`
	res, err := db.Exec(query,
		sql.Named("date", next),
		sql.Named("id", id),
	)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

// AddTask добавляет новую задачу в базу данных.
//
// Возвращает ID созданной задачи.
func AddTask(task *Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`
	res, err := db.Exec(query,
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

// Tasks возвращает список задач, отсортированных по дате.
//
// limit ограничивает количество возвращаемых записей.
func Tasks(limit int) ([]*Task, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT %d", limit))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return processData(rows)
}

// SearchByStringTasks ищет задачи по вхождению строки в title или comment.
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

// SearchByDateTasks возвращает задачи на конкретную дату.
//
// date должен быть в формате YYYYMMDD.
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

// DeleteTask удаляет задачу по ID.
//
// Возвращает ошибку, если задача не найдена.
func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = :id`

	res, err := db.Exec(query, sql.Named("id", id))
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for deleting task`)
	}
	return nil
}

// processData преобразует строки результата SQL-запроса в срез задач.
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
