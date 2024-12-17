package controller

import (
	"SystemMetric/internal/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

// APIGetAlertRecipientsHandler страница получения всех получателей алертов
// @Summary get page
// @Description get alert recipients
// @Tags alert-recipient
// @Accept json
// @Produce json
// @Success 200 {string} string "Get alert recipients is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/alert-recipients [get]
func (s *Server) APIGetAlertRecipientsHandler(w http.ResponseWriter, r *http.Request) {
	alertRecipients, err := s.u.GetAlertRecipients()
	if err != nil {
		http.Error(w, "get alertRecipients is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetAlertRecipientsHandler s.u.GetAlertRecipients",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetAlertRecipientsHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(alertRecipients)
}

// APIGetAlertRecipientHandler страница получения получателя алерта по alertRecipientID
// @Summary get page
// @Description get alert recipient with alertRecipientID
// @Tags alert-recipient
// @Accept json
// @Produce json
// @Param id path int true "AlertRecipient ID"
// @Success 200 {string} string "Get alert recipient is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/alert-recipients/{id} [get]
func (s *Server) APIGetAlertRecipientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	alertRecipientID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "get alertRecipient is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIGetAlertRecipientHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	alertRecipient, err := s.u.GetAlertRecipient(int64(alertRecipientID))
	if err != nil {
		http.Error(w, "get alertRecipient is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetAlertRecipientHandler s.u.GetAlertRecipient",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetAlertRecipientHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(alertRecipient)
}

// APIInsertAlertRecipientHandler страница добавления получателя алерта
// @Summary insert page
// @Description insert alert recipient
// @Tags alert-recipient
// @Accept json
// @Produce json
// @Param alert-recipient body entity.AlertRecipient true "alert_id, user_id"
// @Success 200 {string} string "Insert alert recipient is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/alert-recipients [post]
func (s *Server) APIInsertAlertRecipientHandler(w http.ResponseWriter, r *http.Request) {
	var alertRecipient entity.AlertRecipient
	if err := json.NewDecoder(r.Body).Decode(&alertRecipient); err != nil {
		http.Error(w, "insert alertRecipient is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIInsertAlertRecipientHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	alertRecipientID, err := s.u.InsertAlertRecipient(&alertRecipient)
	if err != nil {
		http.Error(w, "insert alertRecipient is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIInsertAlertRecipientHandler s.u.InsertAlertRecipient",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIInsertAlertRecipientHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(alertRecipientID)
}

// APIDeleteAlertRecipientHandler страница удаления получателя алерта по alertRecipientID
// @Summary delete page
// @Description delete alert recipient with alertRecipientID
// @Tags alert-recipient
// @Accept json
// @Produce json
// @Param id path int true "AlertRecipient ID"
// @Success 200 {string} string "delete alert recipient is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/alert-recipients/{id} [delete]
func (s *Server) APIDeleteAlertRecipientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	alertRecipientID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "delete alertRecipient is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIDeleteAlertRecipientHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	err = s.u.DeleteAlertRecipient(int64(alertRecipientID))
	if err != nil {
		http.Error(w, "delete alertRecipient is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIDeleteAlertRecipientHandler s.u.DeleteAlertRecipient",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIDeleteAlertRecipientHandler", slog.Int("status", http.StatusOK))
}
