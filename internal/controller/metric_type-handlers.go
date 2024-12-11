package controller

import (
	"SystemMetric/internal/entity"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

// APIGetMetricTypesHandler страница получения всех типов метрик
// @Summary get page
// @Description get metric types
// @Tags API
// @Accept json
// @Produce json
// @Success 200 {string} string "Get metric types is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/getMetricTypes [get]
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
// @Tags API
// @Accept json
// @Produce json
// @Param metricTypeID header string true "metricTypeID for getting metricType"
// @Success 200 {string} string "Get metricType is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/getMetricType [get]
func (s *Server) APIGetMetricTypeHandler(w http.ResponseWriter, r *http.Request) {
	metricTypeID, _ := strconv.Atoi(r.Header.Get("metricTypeID"))
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
// @Tags API
// @Accept json
// @Produce json
// @Param metric_type body entity.MetricType true "type_name"
// @Success 200 {string} string "Insert metric type is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/insertMetricType [post]
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
// @Tags API
// @Accept json
// @Produce json
// @Param metricTypeID header string true "metricTypeID for deleting metric type"
// @Success 200 {string} string "delete metric type is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/deleteMetricType [delete]
func (s *Server) APIDeleteMetricTypeHandler(w http.ResponseWriter, r *http.Request) {
	metricTypeID, _ := strconv.Atoi(r.Header.Get("metricTypeID"))
	err := s.u.DeleteMetricType(int64(metricTypeID))
	if err != nil {
		http.Error(w, "delete metricType is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIDeleteMetricTypeHandler s.u.DeleteMetricType",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIDeleteMetricTypeHandler", slog.Int("status", http.StatusOK))
}
