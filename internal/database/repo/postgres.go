package repo

import (
	"SystemMetric/pkg/postgres"
	"context"
)

type PostgresRepo struct {
	db *postgres.Postgres
	//logger *logger.Logger
}

func New(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{
		db: pg,
		//logger: logger,
	}
}

func (r *PostgresRepo) CreateTables() error {
	createTables := []string{`
  create table if not exists users (
	user_id serial primary key,
	username varchar(100) not null,
	email varchar (100) not null unique
);`, `create table if not exists roles (
	role_id serial primary key,
	role_name varchar(100) not null unique
);`, `create table if not exists role_user (
	role_id integer references roles(role_id) not null,
	user_id integer references users(user_id) not null
);`, `create table if not exists metric_types (
	type_id serial primary key,
	type_name varchar(50) not null
);`, `create table if not exists metrics (
	metric_id serial primary key,
	metric_name varchar(50) not null,
	timestamp timestamp with time zone,
	value double precision not null,
	metric_type_id integer references metric_types(type_id) not null
);`, `create table if not exists alerts (
	alert_id serial primary key,
	alert_message text,
	severity integer not null,
	metric_id integer references metrics(metric_id) not null
);`, `
create table if not exists alert_recipients (
	alert_recipient_id serial primary key,
	alert_id integer references alerts(alert_id) not null,
	user_id integer references users(user_id) not null
);`}

	for _, query := range createTables {
		_, err := r.db.Pool.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}
	return nil
}
