package controller

import (
	"SystemMetric/internal/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

// APIGetMetricTypesHandler страница получения всех типов метрик
// @Summary get page
// @Description get metric types
// @Tags metric-type
// @Accept json
// @Produce json
// @Success 200 {string} string "Get metric types is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/metric-types [get]
func (s *Server) APIGetMetricTypesHandler(w http.ResponseWriter, r *http.Request) {
	metricTypes, err := s.u.GetMetricTypes()
	if err != nil {
		http.Error(w, "get metricTypes is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetMetricTypesHandler s.u.GetMetricTypes",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetMetricTypesHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(metricTypes)
}

// APIGetMetricTypeHandler страница получения типа метрики по metricTypeID
// @Summary get page
// @Description get metricType with metricTypeID
// @Tags metric-type
// @Accept json
// @Produce json
// @Param id path int true "MetricType ID"
// @Success 200 {string} string "Get metricType is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/metric-types/{id} [get]
func (s *Server) APIGetMetricTypeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	metricTypeID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "get metricType is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIGetMetricTypeHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	metricType, err := s.u.GetMetricType(int64(metricTypeID))
	if err != nil {
		http.Error(w, "get metricType is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetMetricTypeHandler s.u.GetMetricType",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetMetricTypeHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(metricType)
}

// APIInsertMetricTypeHandler страница добавления типа метрики
// @Summary insert page
// @Description insert metric type
// @Tags metric-type
// @Accept json
// @Produce json
// @Param metric_type body entity.MetricType true "type_name"
// @Success 200 {string} string "Insert metric type is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/metric-types [post]
func (s *Server) APIInsertMetricTypeHandler(w http.ResponseWriter, r *http.Request) {
	var metricType entity.MetricType
	if err := json.NewDecoder(r.Body).Decode(&metricType); err != nil {
		http.Error(w, "insert metricType is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIInsertMetricTypeHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	metricTypeID, err := s.u.InsertMetricType(&metricType)
	if err != nil {
		http.Error(w, "insert metricType is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIInsertMetricTypeHandler s.u.InsertMetricType",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIInsertMetricTypeHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(metricTypeID)
}

// APIDeleteMetricTypeHandler страница удаления типа метрики по metricTypeID
// @Summary delete page
// @Description delete metric type with metricTypeID
// @Tags metric-type
// @Accept json
// @Produce json
// @Param id path int true "MetricType ID"
// @Success 200 {string} string "delete metric type is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/metric-types/{id} [delete]
func (s *Server) APIDeleteMetricTypeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	metricTypeID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "delete metricType is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIDeleteMetricTypeHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	err = s.u.DeleteMetricType(int64(metricTypeID))
	if err != nil {
		http.Error(w, "delete metricType is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIDeleteMetricTypeHandler s.u.DeleteMetricType",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIDeleteMetricTypeHandler", slog.Int("status", http.StatusOK))
}
