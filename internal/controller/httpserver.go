package controller

import (
	"SystemMetric/internal/usecase"
	"SystemMetric/pkg/logger"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
)

type Server struct {
	router    *mux.Router
	u         *usecase.Usecase
	logger    *logger.Logger
	secretKey []byte
}

func New(u *usecase.Usecase, logger *logger.Logger) *Server {
	s := &Server{
		router:    mux.NewRouter(),
		u:         u,
		logger:    logger,
		secretKey: []byte("your_secret_key"),
	}
	//s.router.Use(s.SetHeaders)
	s.router.HandleFunc("/home", s.HomeHandler)
	s.router.HandleFunc("/info", s.InfoHandler)

	//создаем новый маршрутизатор, кот явл подмаршрутизатором основного router, обрабатывающий только url запросы, начинающиеся с api
	apiRouter := s.router.PathPrefix("/api").Subrouter()
	//apiRouter.Use(s.checkToken)
	s.addAPIRouters(apiRouter)
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return s
}

func (s *Server) Run(port string) {
	s.logger.Info("Сервер запущен на http://127.0.0.1:" + port)
	if err := http.ListenAndServe("localhost:"+port, s.router); err != nil {
		s.logger.Error("fatal error", slog.Int("status", http.StatusBadGateway))
		return
	}
}
