package controller

import (
	"SystemMetric/internal/entity"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

// APIGetMetricsHandler страница получения всех метрик
// @Summary get page
// @Description get metrics
// @Tags metric
// @Accept json
// @Produce json
// @Success 200 {string} string "Get metrics is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/metrics [get]
func (s *Server) APIGetMetricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics, err := s.u.GetMetrics()
	if err != nil {
		http.Error(w, "get metrics is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetMetricsHandler s.u.GetMetrics",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetMetricsHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(metrics)
}

// APIGetMetricHandler страница получения метрики по metricID
// @Summary get page
// @Description get metric with metricID
// @Tags metric
// @Accept json
// @Produce json
// @Param id path int true "Metric ID"
// @Success 200 {string} string "Get metric is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/metrics/{id} [get]
func (s *Server) APIGetMetricHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	metricID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "get metric is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIGetMetricHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	metric, err := s.u.GetMetric(int64(metricID))
	if err != nil {
		http.Error(w, "get metric is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetMetricHandler s.u.GetMetric",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetMetricHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(metric)
}

// APIInsertMetricHandler страница добавления метрики
// @Summary insert page
// @Description insert metric
// @Tags metric
// @Accept json
// @Produce json
// @Param metric body entity.Metric true "metric_name, timestamp, value, metric_type_id"
// @Success 200 {string} string "Insert metric is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/metrics [post]
func (s *Server) APIInsertMetricHandler(w http.ResponseWriter, r *http.Request) {
	var metric entity.Metric
	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		http.Error(w, "insert metric is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIInsertInsertMetricHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	metricID, err := s.u.InsertMetric(&metric)
	if err != nil {
		http.Error(w, "insert metric is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIInsertInsertMetricHandler s.u.InsertMetric",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIInsertInsertMetricHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(metricID)
}

// APIDeleteMetricHandler страница удаления метрики по metricID
// @Summary delete page
// @Description delete metric with metricID
// @Tags metric
// @Accept json
// @Produce json
// @Param id path int true "Metric ID"
// @Success 200 {string} string "delete metric is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/metrics/{id} [delete]
func (s *Server) APIDeleteMetricHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	metricID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "delete metric is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIDeleteMetricHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	err = s.u.DeleteMetric(int64(metricID))
	if err != nil {
		http.Error(w, "delete metric is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIDeleteMetricHandler s.u.DeleteMetric",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIDeleteMetricHandler", slog.Int("status", http.StatusOK))
}

// APIUpdateMetricHandler страница обновления значения метрики
// @Summary update page
// @Description update value with metricID
// @Tags metric
// @Accept json
// @Produce json
// @Param metric body entity.Metric true "metric_id, value"
// @Success 200 {string} string "Update metric is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/metrics [put]
func (s *Server) APIUpdateMetricHandler(w http.ResponseWriter, r *http.Request) {
	var metric entity.Metric
	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		http.Error(w, "update metric is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIUpdateMetricHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	validate := validator.New()
	if err := validate.Struct(metric); err != nil {
		http.Error(w, "update metric is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIUpdateMetricHandler validate.Struct",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	err := s.u.UpdateMetric(metric.MetricID, metric.Value)
	if err != nil {
		http.Error(w, "update metric is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIUpdateMetricHandler s.u.UpdateMetric",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIUpdateMetricHandler", slog.Int("status", http.StatusOK))
}
