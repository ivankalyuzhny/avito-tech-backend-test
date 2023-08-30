package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"avito-tech-backend-test/internal/app/service"
	"avito-tech-backend-test/pkg/response"
)

func RegisterSegmentHandlers(router *mux.Router, segmentService *service.SegmentService) {
	handler := NewSegmentHandler(segmentService)

	router.HandleFunc("/segments", handler.CreateSegment).Methods("POST")
	router.HandleFunc("/segments/{slug}", handler.DeleteSegment).Methods("DELETE")
}

type SegmentHandler struct {
	segmentService *service.SegmentService
}

func NewSegmentHandler(segmentService *service.SegmentService) *SegmentHandler {
	return &SegmentHandler{
		segmentService: segmentService,
	}
}

func (h *SegmentHandler) CreateSegment(w http.ResponseWriter, r *http.Request) {
	notEmpty := func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 0
	}

	var input struct {
		Slug string `json:"slug" validate:"notempty"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.WriteError(w, nil, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err = validate.RegisterValidation("notempty", notEmpty)
	if err != nil {
		response.WriteError(w, nil, http.StatusInternalServerError, err)
	}
	err = validate.Struct(input)
	if err != nil {
		response.WriteError(w, nil, http.StatusUnprocessableEntity, err)
	}

	segment, err := h.segmentService.Create(input.Slug)
	if err != nil {
		response.WriteError(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.WriteResponse(w, nil, http.StatusCreated, segment)
}

func (h *SegmentHandler) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	err := h.segmentService.Delete(slug)
	if err != nil {
		response.WriteError(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.WriteResponse(w, nil, http.StatusOK, nil)
}
