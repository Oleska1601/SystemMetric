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
	apiRouter.HandleFunc("/getAlerts", s.APIGetAlertsHandler).Methods("GET")
	apiRouter.HandleFunc("/getAlert", s.APIGetAlertHandler).Methods("GET")
	apiRouter.HandleFunc("/insertAlert", s.APIInsertAlertHandler).Methods("POST")
	apiRouter.HandleFunc("/deleteAlert", s.APIDeleteAlertHandler).Methods("DELETE")

	apiRouter.HandleFunc("/getAlertRecipients", s.APIGetAlertRecipientsHandler).Methods("GET")
	apiRouter.HandleFunc("/getAlertRecipient", s.APIGetAlertRecipientHandler).Methods("GET")
	apiRouter.HandleFunc("/insertAlertRecipient", s.APIInsertAlertRecipientHandler).Methods("POST")
	apiRouter.HandleFunc("/deleteAlertRecipient", s.APIDeleteAlertRecipientHandler).Methods("DELETE")

	apiRouter.HandleFunc("/getMetrics", s.APIGetMetricsHandler).Methods("GET")
	apiRouter.HandleFunc("/getMetric", s.APIGetMetricHandler).Methods("GET")
	apiRouter.HandleFunc("/insertMetric", s.APIInsertMetricHandler).Methods("POST")
	apiRouter.HandleFunc("/deleteMetric", s.APIDeleteMetricHandler).Methods("DELETE")
	apiRouter.HandleFunc("/updateMetric", s.APIUpdateMetricHandler).Methods("PUT")

	apiRouter.HandleFunc("/getMetricTypes", s.APIGetMetricTypesHandler).Methods("GET")
	apiRouter.HandleFunc("/getMetricType", s.APIGetMetricTypeHandler).Methods("GET")
	apiRouter.HandleFunc("/insertMetricType", s.APIInsertMetricTypeHandler).Methods("POST")
	apiRouter.HandleFunc("/deleteMetricType", s.APIDeleteMetricTypeHandler).Methods("DELETE")

	apiRouter.HandleFunc("/getRoles", s.APIGetRolesHandler).Methods("GET")
	apiRouter.HandleFunc("/getRole", s.APIGetRoleHandler).Methods("GET")
	apiRouter.HandleFunc("/insertRole", s.APIInsertRoleHandler).Methods("POST")
	apiRouter.HandleFunc("/deleteRole", s.APIDeleteRoleHandler).Methods("DELETE")

	apiRouter.HandleFunc("/getUsers", s.APIGetUsersHandler).Methods("GET")
	apiRouter.HandleFunc("/getUser", s.APIGetUserHandler).Methods("GET")
	apiRouter.HandleFunc("/insertUser", s.APIInsertUserHandler).Methods("POST")
	apiRouter.HandleFunc("/deleteUser", s.APIDeleteUserHandler).Methods("DELETE")
	apiRouter.HandleFunc("/updateUser", s.APIUpdateUserHandler).Methods("PUT")

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
