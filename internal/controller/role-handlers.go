package controller

import (
	"SystemMetric/internal/entity"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

// APIGetRolesHandler страница получения всех ролей
// @Summary get page
// @Description get roles
// @Tags API
// @Accept json
// @Produce json
// @Success 200 {string} string "Get roles is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/getRoles [get]
func (s *Server) APIGetRolesHandler(w http.ResponseWriter, r *http.Request) {
	roles, err := s.u.GetRoles()
	if err != nil {
		http.Error(w, "get roles is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetRolesHandler s.u.GetRoles",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetRolesHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(roles)
}

// APIGetRoleHandler страница получения роли по roleID
// @Summary get page
// @Description get role with roleID
// @Tags API
// @Accept json
// @Produce json
// @Param roleID header string true "roleID for getting role"
// @Success 200 {string} string "Get role is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/getRole [get]
func (s *Server) APIGetRoleHandler(w http.ResponseWriter, r *http.Request) {
	roleID, _ := strconv.Atoi(r.Header.Get("roleID"))
	role, err := s.u.GetRole(int64(roleID))
	if err != nil {
		http.Error(w, "get role is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetRoleHandler s.u.GetRole",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetRoleHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(role)
}

// APIInsertRoleHandler страница добавления роли
// @Summary insert page
// @Description insert role
// @Tags API
// @Accept json
// @Produce json
// @Param role body entity.Role true "role_name"
// @Success 200 {string} string "Insert role is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/insertRole [post]
func (s *Server) APIInsertRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role entity.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "insert role is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIInsertRoleHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	roleID, err := s.u.InsertRole(&role)
	if err != nil {
		http.Error(w, "insert role is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIInsertRoleHandler s.u.InsertRole",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIInsertRoleHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(roleID)
}

// APIDeleteRoleHandler страница удаления роли по roleID
// @Summary delete page
// @Description delete role with roleID
// @Tags API
// @Accept json
// @Produce json
// @Param roleID header string true "roleID for deleting role"
// @Success 200 {string} string "delete role is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/deleteRole [delete]
func (s *Server) APIDeleteRoleHandler(w http.ResponseWriter, r *http.Request) {
	roleID, _ := strconv.Atoi(r.Header.Get("roleID"))
	err := s.u.DeleteRole(int64(roleID))
	if err != nil {
		http.Error(w, "delete role is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIDeleteRoleHandler s.u.DeleteRole",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIDeleteRoleHandler", slog.Int("status", http.StatusOK))
}
