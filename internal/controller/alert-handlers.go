package controller

import (
	"SystemMetric/internal/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

// APIGetAlertsHandler страница получения всех алертов
// @Summary get page
// @Description get alerts
// @Tags alert
// @Accept json
// @Produce json
// @Success 200 {string} string "Get alerts is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/alerts [get]
func (s *Server) APIGetAlertsHandler(w http.ResponseWriter, r *http.Request) {
	alerts, err := s.u.GetAlerts()
	if err != nil {
		http.Error(w, "get alerts is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetAlertsHandler s.u.GetAlerts",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetAlertsHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(alerts)
}

// APIGetAlertHandler страница получения алерта по alertID
// @Summary get page
// @Description get alert with alertID
// @Tags alert
// @Accept json
// @Produce json
// @Param id path int true "Alert ID"
// @Success 200 {string} string "Get alert is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/alerts/{id} [get]
func (s *Server) APIGetAlertHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	alertID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "get alert is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIGetAlertHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	alert, err := s.u.GetAlert(int64(alertID))
	if err != nil {
		http.Error(w, "get alert is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetAlertHandler s.u.GetAlert",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetAlertHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(alert)
}

// APIInsertAlertHandler страница добавления алерта
// @Summary insert page
// @Description insert alert
// @Tags alert
// @Accept json
// @Produce json
// @Param alert body entity.Alert true "alert_message, severity, metric_id"
// @Success 200 {string} string "Insert alert is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/alerts [post]
func (s *Server) APIInsertAlertHandler(w http.ResponseWriter, r *http.Request) {
	var alert entity.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, "insert alert is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIInsertAlertHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	alertID, err := s.u.InsertAlert(&alert)
	if err != nil {
		http.Error(w, "insert alert is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIInsertAlertHandler s.u.InsertAlert",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIInsertAlertHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(alertID)
}

// APIDeleteAlertHandler страница удаления алерта по alertID
// @Summary delete page
// @Description delete alert
// @Tags alert
// @Accept json
// @Produce json
// @Param id path int true "Alert ID"
// @Success 200 {string} string "delete alert is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/alerts/{id} [delete]
func (s *Server) APIDeleteAlertHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	alertID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "delete alert is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIDeleteAlertHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	err = s.u.DeleteAlert(int64(alertID))
	if err != nil {
		http.Error(w, "delete alert is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIDeleteAlertHandler s.u.DeleteAlert",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIDeleteAlertHandler", slog.Int("status", http.StatusOK))
}
