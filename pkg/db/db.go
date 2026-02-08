package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT,
    repeat VARCHAR(32) NOT NULL DEFAULT ""
);

CREATE INDEX IF NOT EXISTS idx_scheduler_date
ON scheduler(date);
`

const defaultDateFormat = "20060102"

func detect(path string) string {
	if v := os.Getenv("TODO_DBFILE"); v != "" {
		return v
	}
	return path
}

// Init инициализирует соединение с базой данных SQLite
//
// Если файл базы данных не существует, считается что это первый запуск,
// и выполняется установка схемы БД.
//
// Параметры:
//
//	dbFile — путь к файлу базы данных
//
// Возвращает ошибку при невозможности открыть БД или применить схему.
func Init(dbFile string) error {
	file := detect(dbFile)
	_, err := os.Stat(file)
	var install bool
	if err != nil {
		install = true
	}
	db, err = sql.Open("sqlite", file)
	if err != nil {
		return err
	}

	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}

	}
	return nil
}

func CloseDb() {
	db.Close()
}
