package usecase

import (
	"SystemMetric/internal/database/repo"
	"SystemMetric/pkg/logger"
)

type Usecase struct {
	r      *repo.PostgresRepo
	logger *logger.Logger
}

func New(r *repo.PostgresRepo, logger *logger.Logger) *Usecase {
	return &Usecase{
		r:      r,
		logger: logger,
	}
}
