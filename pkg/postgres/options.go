package postgres

import "time"

// объявляем новый тип
type Option func(*Postgres)

// возвращают анонимную функцию, которая принимает указатель на `Postgres` и изменяет соответствующие поля
func MaxPoolSize(size int) Option {
	return func(postgres *Postgres) {
		postgres.maxPoolSize = size
	}
}

func ConnAttempts(attempts int) Option {
	return func(postgres *Postgres) {
		postgres.ConnAttempts = attempts
	}
}

func ConnTimeout(timeout time.Duration) Option {
	return func(postgres *Postgres) {
		postgres.ConnTimeout = timeout
	}
}
