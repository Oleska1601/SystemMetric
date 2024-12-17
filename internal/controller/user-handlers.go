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

// APIGetUsersHandler страница получения всех пользователей
// @Summary get page
// @Description get users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} string "Get users is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users [get]
func (s *Server) APIGetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := s.u.GetUsers()
	if err != nil {
		http.Error(w, "get users is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetUsersHandler s.u.GetUsers",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetUsersHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(users)
}

// APIGetUserHandler страница получения user по userID
// @Summary get page
// @Description get user with userID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "Get user is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users/{id} [get]
func (s *Server) APIGetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	userID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "get user is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIGetUserHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	user, err := s.u.GetUser(int64(userID))
	if err != nil {
		http.Error(w, "get user is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIGetUserHandler s.u.GetUser",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIGetUserHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(user)
}

// APIInsertUserHandler страница добавления пользователя
// @Summary insert page
// @Description insert user
// @Tags user
// @Accept json
// @Produce json
// @Param user body entity.User true "username, email"
// @Success 200 {string} string "Insert user is successful"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users [post]
func (s *Server) APIInsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "insert user is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIInsertUserHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	userID, err := s.u.InsertUser(&user)
	if err != nil {
		http.Error(w, "insert user is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIInsertUserHandler s.u.InsertUser",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIInsertUserHandler", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(userID)
}

// APIDeleteUserHandler страница удаления пользователя по userID
// @Summary delete page
// @Description delete user with userID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "delete user is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users/{id} [delete]
func (s *Server) APIDeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"] // Извлекаем ID из пути
	userID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "delete user is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIDeleteUserHandler strconv.Atoi",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	err = s.u.DeleteUser(int64(userID))
	if err != nil {
		http.Error(w, "delete user is impossible", http.StatusInternalServerError)
		s.logger.Error("controller APIDeleteUserHandler s.u.DeleteUser",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIDeleteUserHandler", slog.Int("status", http.StatusOK))
}

// APIUpdateUserHandler страница обновления почты пользователя
// @Summary update page
// @Description update email with userID
// @Tags user
// @Accept json
// @Produce json
// @Param user body entity.User true "user_id, email"
// @Success 200 {string} string "Update user is successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users [put]
func (s *Server) APIUpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "update user is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIUpdateUserHandler json.NewDecoder(r.Body).Decode",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		http.Error(w, "update user is impossible", http.StatusBadRequest)
		s.logger.Error("controller APIUpdateUserHandler validate.Struct",
			slog.Any("error", err), slog.Int("status", http.StatusBadRequest))
		return
	}
	err := s.u.UpdateUser(user.UserID, user.Email)
	if err != nil {
		http.Error(w, "update user is impossible", http.StatusInternalServerError)
		s.logger.Error("update APIUpdateUserHandler s.u.UpdateUser",
			slog.Any("error", err), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.logger.Info("controller APIUpdateUserHandler", slog.Int("status", http.StatusOK))
}
