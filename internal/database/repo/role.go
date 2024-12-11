package repo

import (
	"SystemMetric/internal/entity"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

func (r *PostgresRepo) GetRoles(ctx context.Context) ([]entity.Role, error) {
	roles := []entity.Role{}
	sql, args, err := r.db.Builder.
		Select("role_id, role_name").
		From("roles").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repo-role GetRoles r.db.Builder: %w", err)
	}
	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repo-role GetRoles r.db.Pool.Query: %w", err)
	}
	for rows.Next() {
		var role entity.Role
		err = rows.Scan(
			&role.RoleID,
			&role.RoleName,
		)
		if err != nil {
			return nil, fmt.Errorf("repo-role GetRoles rows.Scan: %w", err)
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *PostgresRepo) GetRole(ctx context.Context, roleID int64) (entity.Role, error) {
	var role entity.Role
	sql, args, err := r.db.Builder.
		Select("role_id, role_name").
		From("roles").
		Where(squirrel.Eq{"role_id": roleID}).
		ToSql()
	if err != nil {
		return entity.Role{}, fmt.Errorf("repo-role GetRole r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args).Scan(
		&role.RoleID,
		&role.RoleName,
	)
	if err != nil {
		return entity.Role{}, fmt.Errorf("repo-role GetRole r.db.Pool.QueryRow.Scan: %w", err)
	}
	return role, nil
}

func (r *PostgresRepo) InsertRole(ctx context.Context, role *entity.Role) (int64, error) {
	var roleID int64
	sql, args, err := r.db.Builder.
		Insert("roles").
		Columns("role_name").
		Values(role.RoleName).
		Suffix("RETURNING role_id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("repo-role InsertRole r.db.Builder: %w", err)
	}
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&roleID)
	if err != nil {
		return 0, fmt.Errorf("repo-role GetRole r.db.Pool.QueryRow.Scan: %w", err)
	}
	return roleID, nil
}

func (r *PostgresRepo) DeleteRole(ctx context.Context, roleID int64) error {
	sql, args, err := r.db.Builder.
		Delete("roles").
		Where(squirrel.Eq{"role_id": roleID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo-role DeleteRole r.db.Builder: %w", err)
	}
	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo-role DeleteRole r.db.Pool.Exec: %w", err)
	}
	return nil
}
