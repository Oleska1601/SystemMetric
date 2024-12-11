package repo

import (
	"SystemMetric/internal/entity"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

func (r *PostgresRepo) GetAlertRecipients(ctx context.Context) ([]entity.AlertRecipient, error) {
	alertRecipients := []entity.AlertRecipient{}
	sql, args, err := r.db.Builder.
		Select("alert_recipient_id, alert_id, user_id").
		From("alert_recipients").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repo-alert-recipient GetAlertRecipients r.db.Builder: %w", err)
	}
	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repo-alert-recipient GetAlertRecipients r.db.Pool.Query: %w", err)
	}
	for rows.Next() {
		var alertRecipient entity.AlertRecipient
		err = rows.Scan(
			&alertRecipient.AlertRecipientID,
			&alertRecipient.AlertID,
			&alertRecipient.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("repo-alert-recipient GetAlertRecipients rows.Scan: %w", err)
		}
		alertRecipients = append(alertRecipients, alertRecipient)
	}
	return alertRecipients, nil
}

func (r *PostgresRepo) GetAlertRecipient(ctx context.Context, alertRecipientID int64) (entity.AlertRecipient, error) {
	var alertRecipient entity.AlertRecipient
	sql, args, err := r.db.Builder.
		Select("alert_recipient_id, alert_id, user_id").
		From("alert_recipients").
		Where(squirrel.Eq{"alert_recipient_id": alertRecipientID}).
		ToSql()
	if err != nil {
		return entity.AlertRecipient{}, fmt.Errorf("repo-alert-recipient GetAlertRecipient r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&alertRecipient.AlertRecipientID,
		&alertRecipient.AlertID,
		&alertRecipient.UserID,
	)
	if err != nil {
		return entity.AlertRecipient{}, fmt.Errorf("repo-alert-recipient GetAlertRecipient r.db.Pool.QueryRow.Scan: %w", err)
	}
	return alertRecipient, nil
}

func (r *PostgresRepo) InsertAlertRecipient(ctx context.Context, alertRecipient *entity.AlertRecipient) (int64, error) {
	var alertRecipientID int64
	sql, args, err := r.db.Builder.
		Insert("alert_recipients").
		Columns("alert_recipient_id, alert_id, user_id").
		Values(alertRecipient.AlertRecipientID, alertRecipient.AlertID, alertRecipient.UserID).
		Suffix("RETURNING alert_recipient_id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("repo-alert-recipient InsertAlertRecipient r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&alertRecipientID)
	if err != nil {
		return 0, fmt.Errorf("repo-alert-recipient InsertAlertRecipient r.db.Pool.QueryRow.Scan: %w", err)
	}
	return alertRecipientID, nil
}

func (r *PostgresRepo) DeleteAlertRecipient(ctx context.Context, alertRecipientID int64) error {
	sql, args, err := r.db.Builder.
		Delete("alert_recipients").
		Where(squirrel.Eq{"alert_recipient_id": alertRecipientID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo-alert-recipient DeleteAlertRecipient r.db.Builder: %w", err)
	}
	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo-alert-recipient DeleteAlertRecipient r.db.Pool.Exec: %w", err)
	}
	return nil
}
