package usecase

import (
	"SystemMetric/internal/entity"
	"context"
	"errors"
	"log/slog"
)

func (u *Usecase) GetUsers() ([]entity.User, error) {
	ctx, _ := context.WithCancel(context.Background())
	users, err := u.r.GetUsers(ctx)
	if err != nil {
		u.logger.Error("usecase-user GetUsers u.r.GetUsers", slog.Any("error", err))
		return nil, errors.New("internal server error")
	}
	return users, nil
}

func (u *Usecase) GetUser(userID int64) (entity.User, error) {
	ctx, _ := context.WithCancel(context.Background())
	user, err := u.r.GetUser(ctx, userID)
	if err != nil {
		u.logger.Error("usecase-user GetUser u.r.GetUser", slog.Any("error", err))
		return entity.User{}, errors.New("internal server error")
	}
	return user, nil
}

func (u *Usecase) InsertUser(user *entity.User) (int64, error) {
	ctx, _ := context.WithCancel(context.Background())
	userID, err := u.r.InsertUser(ctx, user)
	if err != nil {
		u.logger.Error("usecase-user InsertUser u.r.InsertUser", slog.Any("error", err))
		return 0, errors.New("internal server error")
	}
	return userID, nil
}

func (u *Usecase) DeleteUser(userID int64) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.r.DeleteUser(ctx, userID)
	if err != nil {
		u.logger.Error("usecase-user DeleteUser u.r.DeleteUser", slog.Any("error", err))
		return errors.New("internal server error")
	}
	return nil
}

func (u *Usecase) UpdateUser(userID int64, newEmail string) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.r.UpdateUser(ctx, userID, newEmail)
	if err != nil {
		u.logger.Error("usecase-user UpdateUser u.r.UpdateUser", slog.Any("error", err))
		return errors.New("internal server error")
	}
	return nil
}
