package repo

import (
	"SystemMetric/internal/entity"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

func (r *PostgresRepo) GetAlerts(ctx context.Context) ([]entity.Alert, error) {
	sql, args, err := r.db.Builder.
		Select("alert_id, alert_message, severity, metric_id").
		From("alerts").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repo-alert GetAlerts r.db.Builder: %w", err)
	}
	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repo-alert GetAlerts r.db.Pool.Query: %w", err)
	}
	alerts := []entity.Alert{}
	for rows.Next() {
		var alert entity.Alert
		err = rows.Scan(
			&alert.AlertID,
			&alert.AlertMessage,
			&alert.Severity,
			&alert.MetricID,
		)
		if err != nil {
			return nil, fmt.Errorf("repo-alert GetAlerts rows.Scan: %w", err)
		}
		alerts = append(alerts, alert)
	}
	return alerts, nil
}

func (r *PostgresRepo) GetAlert(ctx context.Context, alertID int64) (entity.Alert, error) {
	var alert entity.Alert
	sql, args, err := r.db.Builder.
		Select("alert_id", "alert_message, severity, metric_id").
		From("alerts").
		Where(squirrel.Eq{"alert_id": alertID}).
		ToSql()
	if err != nil {
		return entity.Alert{}, fmt.Errorf("repo-alert GetAlert r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&alert.AlertID,
		&alert.AlertMessage,
		&alert.Severity,
		&alert.MetricID,
	)
	if err != nil {
		return entity.Alert{}, fmt.Errorf("repo-alert GetAlert r.db.Pool.QueryRow.Scan: %w", err)
	}
	return alert, nil
}

func (r *PostgresRepo) InsertAlert(ctx context.Context, alert *entity.Alert) (int64, error) {
	var alertID int64
	sql, args, err := r.db.Builder.
		Insert("alerts").
		Columns("alert_message, severity, metric_id").
		Values(alert.AlertMessage, alert.Severity, alert.MetricID).
		Suffix("RETURNING alert_id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("repo-alert InsertAlert r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&alertID)
	if err != nil {
		return 0, fmt.Errorf("repo-alert InsertAlert r.db.Pool.QueryRow.Scan: %w", err)
	}
	return alertID, nil
}

func (r *PostgresRepo) DeleteAlert(ctx context.Context, alertID int64) error {
	sql, args, err := r.db.Builder.Delete("alerts").
		Where(squirrel.Eq{"alert_id": alertID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo-alert DeleteAlert r.db.Builder: %w", err)
	}
	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo-alert r.db.Pool.Exec: %w", err)
	}
	return nil
}
