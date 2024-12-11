package usecase

import (
	"SystemMetric/internal/entity"
	"context"
	"errors"
	"log/slog"
)

func (u *Usecase) GetAlertRecipients() ([]entity.AlertRecipient, error) {
	ctx, _ := context.WithCancel(context.Background())
	alertRecipients, err := u.r.GetAlertRecipients(ctx)
	if err != nil {
		u.logger.Error("usecase-alert-recipient GetAlertRecipients u.r.GetAlertRecipients", slog.Any("error", err))
		return nil, errors.New("internal server error")
	}

	return alertRecipients, nil
}

func (u *Usecase) GetAlertRecipient(alertRecipientID int64) (entity.AlertRecipient, error) {
	ctx, _ := context.WithCancel(context.Background())
	alertRecipient, err := u.r.GetAlertRecipient(ctx, alertRecipientID)
	if err != nil {
		u.logger.Error("usecase-alert-recipient GetAlertRecipient u.r.GetAlertRecipient", slog.Any("error", err))
		return entity.AlertRecipient{}, errors.New("internal server error")
	}
	return alertRecipient, nil
}

func (u *Usecase) InsertAlertRecipient(alertRecipient *entity.AlertRecipient) (int64, error) {
	ctx, _ := context.WithCancel(context.Background())
	alertRecipientID, err := u.r.InsertAlertRecipient(ctx, alertRecipient)
	if err != nil {
		u.logger.Error("usecase-alert-recipient InsertAlertRecipient u.r.InsertAlertRecipient", slog.Any("error", err))
		return 0, errors.New("internal server error")
	}
	return alertRecipientID, nil
}

func (u *Usecase) DeleteAlertRecipient(alertRecipientID int64) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.r.DeleteAlertRecipient(ctx, alertRecipientID)
	if err != nil {
		u.logger.Error("usecase-alert-recipient DeleteAlertRecipient u.r.DeleteAlertRecipient", slog.Any("error", err))
		return errors.New("internal server error")
	}
	return nil
}
