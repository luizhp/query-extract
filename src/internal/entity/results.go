package entity

import "time"

type Results struct {
	Columns    []string
	Rows       []map[string]interface{}
	StartedAt  time.Time
	FinishedAt time.Time
}

func NewResults(columns []string, rows []map[string]interface{}, startedAt, finishedAt time.Time) *Results {
	return &Results{
		Columns:    columns,
		Rows:       rows,
		StartedAt:  startedAt,
		FinishedAt: finishedAt,
	}
}

func (r *Results) GetColumns() []string {
	return r.Columns
}

func (r *Results) GetRows() []map[string]interface{} {
	return r.Rows
}

func (r *Results) GetTotalRows() int {
	return len(r.Rows)
}

func (r *Results) GetStartedAt() time.Time {
	return r.StartedAt
}

func (r *Results) GetFinishedAt() time.Time {
	return r.FinishedAt
}

func (r *Results) GetDuration() time.Duration {
	return r.FinishedAt.Sub(r.StartedAt)
}
