package main

import (
	"SystemMetric/config"
	_ "SystemMetric/docs"
	"SystemMetric/internal/controller"
	"SystemMetric/internal/database/repo"
	"SystemMetric/internal/usecase"
	"SystemMetric/pkg/logger"
	"SystemMetric/pkg/postgres"
	"log/slog"
)

// @title SystemMetric API
// @version 1.0
// @description This is a server for work with metrics
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("main config.New", slog.Any("error", err))
		return
	}
	l := logger.New(cfg.Logger.Level)
	pg, err := postgres.New(l, cfg.Postgres.PgUrl)
	defer pg.CLose()
	if err != nil {
		l.Error("main postgres.New", slog.Any("error", err))
		return
	}
	db := repo.New(pg)
	if err = db.CreateTables(); err != nil {
		l.Error("db.CreateTables", slog.Any("error", err))
		return
	}
	l.Info("connection to db")
	u := usecase.New(db, l)
	s := controller.New(u, l)
	s.Run(cfg.HTTP.Port)
}
