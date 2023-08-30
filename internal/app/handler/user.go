package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"avito-tech-backend-test/internal/app/service"
	"avito-tech-backend-test/pkg/response"
)

func RegisterUserHandlers(router *mux.Router, userService *service.UserService) {
	handler := NewUserHandler(userService)

	router.HandleFunc("/users/{id}/segments", handler.FindUserSegments).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}/segments", handler.EditUserSegments).Methods(http.MethodPut)
}

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) FindUserSegments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response.WriteResponse(
			w,
			map[string]string{"Content-Type": "application/json"},
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	userSegments, err := h.userService.FindSegmentsByUserID(userID)
	if err != nil {
		response.WriteResponse(
			w,
			map[string]string{"Content-Type": "application/json"},
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	response.WriteResponse(
		w,
		map[string]string{"Content-Type": "application/json"},
		http.StatusOK,
		userSegments,
	)
}

func (h *UserHandler) EditUserSegments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var input struct {
		SegmentsAdd []string `json:"segments_add" validate:"required"`
		SegmentsDel []string `json:"segments_del" validate:"required"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.WriteResponse(
			w,
			map[string]string{"Content-Type": "application/json"},
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		response.WriteResponse(
			w,
			map[string]string{"Content-Type": "application/json"},
			http.StatusUnprocessableEntity,
			err.Error(),
		)
	}

	err = h.userService.EditSegments(userID, input.SegmentsAdd, input.SegmentsDel)
	if err != nil {
		response.WriteResponse(
			w,
			map[string]string{"Content-Type": "application/json"},
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	response.WriteResponse(
		w,
		map[string]string{"Content-Type": "application/json"},
		http.StatusOK,
		nil,
	)
}
