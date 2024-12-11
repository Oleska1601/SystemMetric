package entity

type AlertRecipient struct {
	AlertRecipientID int64 `json:"alert_recipient_id" example:"1"`
	AlertID          int64 `json:"alert_id" example:"1"`
	UserID           int64 `json:"user_id" example:"1"`
}
