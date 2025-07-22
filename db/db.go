package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Task represents a task in the tracker.
func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "tracker.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InsertLog inserts a new log entry into the database.
// It takes a task name as an argument and returns an error if the insertion fails.
func InsertLog(db *sql.DB, category string, value string) error {
	_, err := db.Exec("INSERT INTO logs (category, value) VALUES (?, ?)", category, value)
	return err
}

// CloseDB closes the database connection.
func CloseDB(db *sql.DB) error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// GetLogs retrieves all logs from the database.
// It returns a slice of strings containing the task names.
func GetLogs(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT task_name FROM logs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []string
	for rows.Next() {
		var taskName string
		if err := rows.Scan(&taskName); err != nil {
			return nil, err
		}
		logs = append(logs, taskName)
	}
	return logs, nil
}

// Task represents a task in the tracker.
// It contains the task name and other relevant fields.
// This struct is used to define the structure of tasks in the database.
