package models

type ExecuteRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
	Input    string `json:"input,omitempty"`
}

type ExecuteJob struct {
	JobID    string
	Language string
	Code     string
	Input    string
}

type ExecuteResult struct {
	JobID    string
	Stdout   string
	Stderr   string
	Error    string
	ExitCode int
}
