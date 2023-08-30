package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"avito-tech-backend-test/internal/app/service"
	"avito-tech-backend-test/pkg/response"
)

func RegisterUserHandlers(router *mux.Router, userService *service.UserService) {
	handler := NewUserHandler(userService)

	router.HandleFunc("/users/{id}/segments", handler.FindUserSegments).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}/segments", handler.UpdateUserSegments).Methods(http.MethodPut)
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
		response.WriteError(w, nil, http.StatusBadRequest, err)
		return
	}

	userSegments, err := h.userService.FindUserSegments(userID)
	if err != nil {
		response.WriteError(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.WriteResponse(w, nil, http.StatusOK, userSegments)
}

func (h *UserHandler) UpdateUserSegments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response.WriteError(w, nil, http.StatusBadRequest, err)
		return
	}

	var input struct {
		SegmentsAdd []string `json:"segments_add" validate:"required"`
		SegmentsDel []string `json:"segments_del" validate:"required"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.WriteError(w, nil, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		response.WriteError(w, nil, http.StatusUnprocessableEntity, err)
	}

	err = h.userService.UpdateUserSegments(userID, input.SegmentsAdd, input.SegmentsDel)
	if err != nil {
		response.WriteError(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.WriteResponse(w, nil, http.StatusOK, nil)
}
