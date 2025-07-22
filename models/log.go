package models

import "time"

type Log struct {
	TaskName  string    `json:"task_name"`
	Timestamp time.Time `json:"timestamp"`
	Category  string    `json:"category"`
	Value     string    `json:"value"`
}
