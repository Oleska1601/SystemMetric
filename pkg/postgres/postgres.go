package postgres

import (
	"SystemMetric/pkg/logger"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"log/slog"
	"time"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize  int
	ConnAttempts int
	ConnTimeout  time.Duration
	Builder      squirrel.StatementBuilderType
	Pool         *pgxpool.Pool
}

func New(logger *logger.Logger, url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		ConnAttempts: _defaultConnAttempts,
		ConnTimeout:  _defaultConnTimeout,
	}
	for _, opt := range opts {
		opt(pg)
	}

	//$1, $2...
	//настраиваем Builder запросов, устанавливаем формат Placeholder-ов (заменителей значений) $1, $2... или '?'
	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	//создает конфигурацию для пула соединений - парсит url и создает pgxpool.COnfig
	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	for pg.ConnAttempts > 0 {
		//создает пул соединений на основе предоставленной конфигурации
		//Указатель на созданный пул соединений
		pg.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}
		logger.Info("Postgres is trying to connect", slog.Int("attempts left", pg.ConnAttempts))
		time.Sleep(pg.ConnTimeout)
		pg.ConnAttempts--

	}
	if err != nil {
		return nil, fmt.Errorf("postgres - New - connAttempts == 0: %w", err)
	}
	return pg, nil
}

func (pg *Postgres) CLose() {
	if pg.Pool != nil {
		pg.Pool.Close()
	}
}
