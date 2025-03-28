package entity

import (
	"time"
)

type Result struct {
	columns    []Column
	rows       []map[string]interface{}
	startedAt  time.Time
	finishedAt time.Time
}

func NewResult(columns []Column, rows []map[string]interface{}, startedAt time.Time, finishedAt time.Time) *Result {
	return &Result{
		columns:    columns,
		rows:       rows,
		startedAt:  startedAt,
		finishedAt: finishedAt,
	}
}

func (r *Result) GetColumnsMetadata() []Column {
	return r.columns
}

func (r *Result) GetColumnsName() []string {
	var columns []string
	for _, c := range r.columns {
		columns = append(columns, c.GetName())
	}
	return columns
}

func (r *Result) GetRows() []map[string]interface{} {
	return r.rows
}

func (r *Result) GetTotalRows() int {
	return len(r.rows)
}

func (r *Result) GetStartedAt() time.Time {
	return r.startedAt
}

func (r *Result) GetFinishedAt() time.Time {
	return r.finishedAt
}

func (r *Result) GetDuration() time.Duration {
	return r.finishedAt.Sub(r.startedAt)
}
