package models

type Task struct {
	Id          int64  `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	IsFinished  bool   `json:"is_finished"`
	Created     string `json:"created,omitempty"`
}
