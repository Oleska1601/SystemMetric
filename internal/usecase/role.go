package usecase

import (
	"SystemMetric/internal/entity"
	"context"
	"errors"
	"log/slog"
)

func (u *Usecase) GetRoles() ([]entity.Role, error) {
	ctx, _ := context.WithCancel(context.Background())
	roles, err := u.r.GetRoles(ctx)
	if err != nil {
		u.logger.Error("usecase-role GetRoles u.r.GetRoles", slog.Any("error", err))
		return nil, errors.New("internal server error")
	}
	return roles, nil
}

func (u *Usecase) GetRole(roleID int64) (entity.Role, error) {
	ctx, _ := context.WithCancel(context.Background())
	role, err := u.r.GetRole(ctx, roleID)
	if err != nil {
		u.logger.Error("usecase-role GetRole u.r.GetRole", slog.Any("error", err))
		return entity.Role{}, errors.New("internal server error")
	}
	return role, nil
}

func (u *Usecase) InsertRole(role *entity.Role) (int64, error) {
	ctx, _ := context.WithCancel(context.Background())
	roleID, err := u.r.InsertRole(ctx, role)
	if err != nil {
		u.logger.Error("usecase-role InsertRole u.r.InsertRole", slog.Any("error", err))
		return 0, errors.New("internal server error")
	}
	return roleID, nil
}

func (u *Usecase) DeleteRole(roleID int64) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.r.DeleteRole(ctx, roleID)
	if err != nil {
		u.logger.Error("usecase-role DeleteRole u.r.DeleteRole", slog.Any("error", err))
		return errors.New("internal server error")
	}
	return nil
}
