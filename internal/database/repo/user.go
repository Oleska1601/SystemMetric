package repo

import (
	"SystemMetric/internal/entity"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

func (r *PostgresRepo) GetUsers(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{}
	sql, args, err := r.db.Builder.
		Select("user_id, username, email").
		From("users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repo-user GetUsers r.db.Builder: %w", err)
	}
	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repo-user GetUsers r.db.Pool.Query: %w", err)
	}
	for rows.Next() {
		var user entity.User
		err = rows.Scan(
			&user.UserID,
			&user.Username,
			&user.Email,
		)
		if err != nil {
			return nil, fmt.Errorf("repo-user GetUsers rows.Scan: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *PostgresRepo) GetUser(ctx context.Context, userID int64) (entity.User, error) {
	var user entity.User
	sql, args, err := r.db.Builder.
		Select("user_id, username, email").
		From("users").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("repo-user GetUser r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
	)
	if err != nil {
		return entity.User{}, fmt.Errorf("repo-user GetUser r.db.Pool.QueryRow.Scan: %w", err)
	}
	return user, nil
}

func (r *PostgresRepo) InsertUser(ctx context.Context, user *entity.User) (int64, error) {
	var userID int64
	sql, args, err := r.db.Builder.
		Insert("users").
		Columns("username, email").
		Values(user.Username, user.Email).
		Suffix("RETURNING user_id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("repo-user InsertUser r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("repo-user InsertUser r.db.Pool.QueryRow.Scan: %w", err)
	}
	return userID, nil
}

func (r *PostgresRepo) DeleteUser(ctx context.Context, userID int64) error {
	sql, args, err := r.db.Builder.
		Delete("users").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo-user DeleteUser r.db.Builder: %w", err)
	}
	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo-user DeleteUser r.db.Pool.Exec: %w", err)
	}
	return nil
}

func (r *PostgresRepo) UpdateUser(ctx context.Context, userID int64, newEmail string) error {
	sql, args, err := r.db.Builder.
		Update("users").
		Set("email", newEmail).
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo-user UpdateUser r.db.Builder: %w", err)
	}
	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo-user UpdateUser r.db.Pool.Exec: %w", err)
	}
	return nil
}
