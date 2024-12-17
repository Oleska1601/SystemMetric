package entity

import "time"

type Metric struct {
	MetricID     int64     `json:"metric_id" example:"1" validate:"required"`
	MetricName   string    `json:"metric_name" example:"metric name"`
	Timestamp    time.Time `json:"timestamp" example:"2006-01-02T15:04:05Z"`
	Value        float64   `json:"value" example:"0.5" validate:"required"`
	MetricTypeID int64     `json:"metric_type_id" example:"1"`
}
