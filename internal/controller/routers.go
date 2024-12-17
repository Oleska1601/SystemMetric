package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) addAPIRouters(apiRouter *mux.Router) {
	s.addAlertsRouters(apiRouter)
	s.addAlertsRecipientsRouters(apiRouter)
	s.addMetricsRouters(apiRouter)
	s.addMetricTypesRouters(apiRouter)
	s.addRolesRouters(apiRouter)
	s.addUsersRouters(apiRouter)
}

func (s *Server) addAlertsRouters(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/alerts", s.APIGetAlertsHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/alerts/{id}", s.APIGetAlertHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/alerts", s.APIInsertAlertHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/alerts/{id}", s.APIDeleteAlertHandler).Methods(http.MethodDelete)
}

func (s *Server) addAlertsRecipientsRouters(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/alert-recipients", s.APIGetAlertRecipientsHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/alert-recipients/{id}", s.APIGetAlertRecipientHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/alert-recipients", s.APIInsertAlertRecipientHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/alert-recipients/{id}", s.APIDeleteAlertRecipientHandler).Methods(http.MethodDelete)
}

func (s *Server) addMetricsRouters(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/metrics", s.APIGetMetricsHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/metrics/{id}", s.APIGetMetricHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/metrics", s.APIInsertMetricHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/metrics/{id}", s.APIDeleteMetricHandler).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/metrics", s.APIUpdateMetricHandler).Methods(http.MethodPut)
}

func (s *Server) addMetricTypesRouters(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/metric-types", s.APIGetMetricTypesHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/metric-types/{id}", s.APIGetMetricTypeHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/metric-types", s.APIInsertMetricTypeHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/metric-types/{id}", s.APIDeleteMetricTypeHandler).Methods(http.MethodDelete)
}

func (s *Server) addRolesRouters(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/roles", s.APIGetRolesHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/roles/{id}", s.APIGetRoleHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/roles", s.APIInsertRoleHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/roles/{id}", s.APIDeleteRoleHandler).Methods(http.MethodDelete)
}

func (s *Server) addUsersRouters(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/users", s.APIGetUsersHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/users/{id}", s.APIGetUserHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc("/users", s.APIInsertUserHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/users/{id}", s.APIDeleteUserHandler).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/users", s.APIUpdateUserHandler).Methods(http.MethodPut)
}
