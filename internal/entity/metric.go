package entity

type Metric struct {
	MetricID     int64   `json:"metric_id" example:"1"`
	MetricName   string  `json:"metric_name" example:"metric name"`
	Timestamp    string  `json:"timestamp" example:"1993-12-01 10:03:44.000 +0300"`
	Value        float64 `json:"value" example:"0.5"`
	MetricTypeID int64   `json:"metric_type_id" example:"1"`
}
