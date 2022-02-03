package models

type Task struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	IsFinished  bool   `json:"is_finished"`
	Created     string `json:"created"`
	Finished    string `json:"finished,omitempty"`
}
