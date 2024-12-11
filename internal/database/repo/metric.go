package repo

import (
	"SystemMetric/internal/entity"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

func (r *PostgresRepo) GetMetrics(ctx context.Context) ([]entity.Metric, error) {
	metrics := []entity.Metric{}
	sql, args, err := r.db.Builder.
		Select("metric_id, metric_name, timestamp, value, metric_type_id").
		From("metrics").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repo-metric GetMetrics r.db.Builder: %w", err)
	}
	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repo-metric r.db.Pool.Query: %w", err)
	}
	for rows.Next() {
		var metric entity.Metric
		err = rows.Scan(
			&metric.MetricID,
			&metric.MetricName,
			&metric.Timestamp,
			&metric.Value,
			&metric.MetricTypeID,
		)
		if err != nil {
			return nil, fmt.Errorf("repo-metric GetMetrics rows.Scan: %w", err)
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

func (r *PostgresRepo) GetMetric(ctx context.Context, metricID int64) (entity.Metric, error) {
	var metric entity.Metric
	sql, args, err := r.db.Builder.
		Select("metric_id, metric_name, timestamp, value, metric_type_id").
		From("metrics").
		Where(squirrel.Eq{"metric_id": metricID}).
		ToSql()
	if err != nil {
		return entity.Metric{}, fmt.Errorf("repo-metric GetMetric r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&metric.MetricID,
		&metric.MetricName,
		&metric.Timestamp,
		&metric.Value,
		&metric.MetricTypeID,
	)
	if err != nil {
		return entity.Metric{}, fmt.Errorf("repo-metric GetMetric r.db.QueryRow.Scan: %w", err)
	}
	return metric, nil
}

func (r *PostgresRepo) InsertMetric(ctx context.Context, metric *entity.Metric) (int64, error) {
	var metricID int64
	sql, args, err := r.db.Builder.
		Insert("metrics").
		Columns("metric_name, timestamp, value, metric_type_id").
		Values(metric.MetricName, metric.Timestamp, metric.Value, metric.MetricTypeID).
		Suffix("RETURNING metric_id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("repo-metric InsertMetric r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&metricID)
	if err != nil {
		return 0, fmt.Errorf("repo-metric InsertMetric r.db.Pool.QueryRow.Scan: %w", err)
	}
	return metricID, err
}

func (r *PostgresRepo) DeleteMetric(ctx context.Context, metricID int64) error {
	sql, args, err := r.db.Builder.
		Delete("metrics").
		Where(squirrel.Eq{"metric_id": metricID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo-metric DeleteMetric r.db.Builder: %w", err)
	}
	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo-metric DeleteMetric r.db.Pool.Exec: %w", err)
	}
	return nil
}

func (r *PostgresRepo) UpdateMetric(ctx context.Context, metricID int64, newValue float64) error {
	sql, args, err := r.db.Builder.
		Update("metrics").
		Set("value", newValue).
		Where(squirrel.Eq{"metric_id": metricID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo-metric UpdateMetric r.db.Builder: %w", err)
	}
	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo-metric UpdateMetric r.db.Pool.Exec: %w", err)
	}
	return nil
}
