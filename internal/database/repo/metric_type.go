package repo

import (
	"SystemMetric/internal/entity"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

func (r *PostgresRepo) GetMetricTypes(ctx context.Context) ([]entity.MetricType, error) {
	metricTypes := []entity.MetricType{}
	sql, args, err := r.db.Builder.
		Select("type_id, type_name").
		From("metric_types").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repo-metric_type GetMetricTypes r.db.Builder: %w", err)
	}
	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repo-metric_type GetMetricTypes r.db.Pool.Query: %w", err)
	}
	for rows.Next() {
		var metricType entity.MetricType
		err = rows.Scan(
			&metricType.TypeID,
			&metricType.TypeName,
		)
		if err != nil {
			return nil, fmt.Errorf("repo-metric_type GetMetricTypes rows.Scan: %w", err)
		}
		metricTypes = append(metricTypes, metricType)
	}
	return metricTypes, nil
}

func (r *PostgresRepo) GetMetricType(ctx context.Context, typeID int64) (entity.MetricType, error) {
	var metricType entity.MetricType
	sql, args, err := r.db.Builder.
		Select("type_id, type_name").
		From("metric_types").
		Where(squirrel.Eq{"type_id": typeID}).
		ToSql()
	if err != nil {
		return entity.MetricType{}, fmt.Errorf("repo-metric_type GetMetricType r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&metricType.TypeID,
		&metricType.TypeName,
	)
	if err != nil {
		return entity.MetricType{}, fmt.Errorf("repo-metric_type GetMetricType r.db.Pool.QueryRow.Scan: %w", err)
	}
	return metricType, err
}

func (r *PostgresRepo) InsertMetricType(ctx context.Context, metricType *entity.MetricType) (int64, error) {
	var typeID int64
	sql, args, err := r.db.Builder.
		Insert("metric_types").
		Columns("type_name").
		Values(metricType.TypeName).
		Suffix("RETURNING type_id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("repo-metric_type InsertMetricType r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&typeID)
	if err != nil {
		return 0, fmt.Errorf("repo-metric_type InsertMetricType r.db.Pool.QueryRow.Scan: %w", err)
	}
	return typeID, nil
}

func (r *PostgresRepo) DeleteMetricType(ctx context.Context, typeID int64) error {
	sql, args, err := r.db.Builder.
		Delete("metric_types").
		Where(squirrel.Eq{"type_id": typeID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo-metric_type DeleteMetricType r.db.Builder: %w", err)
	}
	_, err = r.db.Pool.Exec(ctx, sql, args)
	if err != nil {
		return fmt.Errorf("repo-metric_type DeleteMetricType r.db.Pool.Exec: %w", err)
	}
	return nil
}
