package usecase

import (
	"SystemMetric/internal/entity"
	"context"
	"errors"
	"log/slog"
)

func (u *Usecase) GetMetricTypes() ([]entity.MetricType, error) {
	ctx, _ := context.WithCancel(context.Background())
	metricTypes, err := u.r.GetMetricTypes(ctx)
	if err != nil {
		u.logger.Error("usecase-metric_type GetMetricTypes u.r.GetMetricTypes", slog.Any("error", err))
		return nil, errors.New("internal server error")
	}
	return metricTypes, nil
}

func (u *Usecase) GetMetricType(typeID int64) (entity.MetricType, error) {
	ctx, _ := context.WithCancel(context.Background())
	metricType, err := u.r.GetMetricType(ctx, typeID)
	if err != nil {
		u.logger.Error("usecase-metric_type GetMetricType u.r.GetMetricType", slog.Any("error", err))
		return entity.MetricType{}, errors.New("internal server error")
	}
	return metricType, nil
}

func (u *Usecase) InsertMetricType(metricType *entity.MetricType) (int64, error) {
	ctx, _ := context.WithCancel(context.Background())
	metricTypeID, err := u.r.InsertMetricType(ctx, metricType)
	if err != nil {
		u.logger.Error("usecase-metric_type InsertMetricType u.r.InsertMetricType", slog.Any("error", err))
		return 0, errors.New("internal server error")
	}
	return metricTypeID, nil
}

func (u *Usecase) DeleteMetricType(typeID int64) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.r.DeleteMetricType(ctx, typeID)
	if err != nil {
		u.logger.Error("usecase-metric_type DeleteMetricType u.r.DeleteMetricType", slog.Any("error",err))
		return errors.New("internal server error")
	}
	return nil
}
