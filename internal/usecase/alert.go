package usecase

import (
	"SystemMetric/internal/entity"
	"context"
	"errors"
	"log/slog"
)

func (u *Usecase) GetAlerts() ([]entity.Alert, error) {
	ctx, _ := context.WithCancel(context.Background())
	alerts, err := u.r.GetAlerts(ctx)
	if err != nil {
		u.logger.Error("usecase-alert GetAlerts u.r.GetAlerts", slog.Any("error", err))
		return nil, errors.New("internal server error")
	}
	return alerts, nil
}

func (u *Usecase) GetAlert(alertID int64) (entity.Alert, error) {
	ctx, _ := context.WithCancel(context.Background())
	alert, err := u.r.GetAlert(ctx, alertID)
	if err != nil {
		u.logger.Error("usecase-alert GetAlert u.r.GetAlert", slog.Any("error", err))
		return entity.Alert{}, errors.New("internal server error")
	}
	return alert, nil
}

func (u *Usecase) InsertAlert(alert *entity.Alert) (int64, error) {
	ctx, _ := context.WithCancel(context.Background())
	alertID, err := u.r.InsertAlert(ctx, alert)
	if err != nil {
		u.logger.Error("usecase-alert InsertAlert u.r.InsertAlert", slog.Any("error", err))
		return 0, errors.New("internal server error")
	}
	return alertID, nil
}

func (u *Usecase) DeleteAlert(alertID int64) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.r.DeleteAlert(ctx, alertID)
	if err != nil {
		u.logger.Error("usecase-alert DeleteAlert u.r.DeleteAlert", slog.Any("error", err))
		return errors.New("internal server error")
	}
	return nil
}
