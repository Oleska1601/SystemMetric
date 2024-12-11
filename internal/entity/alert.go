package entity

type Alert struct {
	AlertID      int64  `json:"alert_id" example:"1"`
	AlertMessage string `json:"alert_message" example:"alert message"`
	Severity     int    `json:"severity" example:"1"`
	MetricID     int64  `json:"metric_id" example:"1"`
}
