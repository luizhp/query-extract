package entity

import (
	"time"
)

type Result struct {
	Columns    []string
	Rows       []map[string]interface{}
	TotalRows  int
	StartedAt  time.Time
	FinishedAt time.Time
}

func NewResult(columns []string, rows []map[string]interface{}, startedAt time.Time, finishedAt time.Time) *Result {
	return &Result{
		Columns:    columns,
		Rows:       rows,
		TotalRows:  len(rows),
		StartedAt:  startedAt,
		FinishedAt: finishedAt,
	}
}

func (r *Result) GetColumns() []string {
	return r.Columns
}

func (r *Result) GetRows() []map[string]interface{} {
	return r.Rows
}

func (r *Result) GetTotalRows() int {
	return len(r.Rows)
}

func (r *Result) GetStartedAt() time.Time {
	return r.StartedAt
}

func (r *Result) GetFinishedAt() time.Time {
	return r.FinishedAt
}

func (r *Result) GetDuration() time.Duration {
	return r.FinishedAt.Sub(r.StartedAt)
}
