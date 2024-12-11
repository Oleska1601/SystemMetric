package usecase

import (
	"SystemMetric/internal/entity"
	"context"
	"errors"
	"log/slog"
)

func (u *Usecase) GetMetrics() ([]entity.Metric, error) {
	ctx, _ := context.WithCancel(context.Background())
	metrics, err := u.r.GetMetrics(ctx)
	if err != nil {
		u.logger.Error("usecase-metric GetMetrics u.r.GetMetrics", slog.Any("error", err))
		return nil, errors.New("internal server error")
	}
	return metrics, nil
}

func (u *Usecase) GetMetric(metricID int64) (entity.Metric, error) {
	ctx, _ := context.WithCancel(context.Background())
	metric, err := u.r.GetMetric(ctx, metricID)
	if err != nil {
		u.logger.Error("usecase-metric GetMetric u.r.GetMetric", slog.Any("error", err))
		return entity.Metric{}, errors.New("internal server error")
	}
	return metric, nil
}

func (u *Usecase) InsertMetric(metric *entity.Metric) (int64, error) {
	ctx, _ := context.WithCancel(context.Background())
	metricID, err := u.r.InsertMetric(ctx, metric)
	if err != nil {
		u.logger.Error("usecase-metric InsertMetric u.r.InsertMetric", slog.Any("error", err))
		return 0, errors.New("internal server error")
	}
	return metricID, nil
}

func (u *Usecase) DeleteMetric(metricID int64) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.r.DeleteMetric(ctx, metricID)
	if err != nil {
		u.logger.Error("usecase-metric DeleteMetric u.r.DeleteMetric", slog.Any("error", err))
		return errors.New("internal server error")
	}
	return nil
}

func (u *Usecase) UpdateMetric(metricID int64, newValue float64) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.r.UpdateMetric(ctx, metricID, newValue)
	if err != nil {
		u.logger.Error("usecase-metric UpdateMetric u.r.UpdateMetric", slog.Any("error", err))
		return errors.New("internal server error")
	}
	return nil
}
