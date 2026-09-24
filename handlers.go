package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type LessonHandler struct {
	store *LessonStore
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (h *LessonHandler) List(w http.ResponseWriter, r *http.Request) {
	lessons, err := h.store.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusOK, lessons)
}

func (h *LessonHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	lesson, err := h.store.Get(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "lesson not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusOK, lesson)
}

func (h *LessonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		GroupID   int    `json:"group_id"`
		TeacherID int    `json:"teacher_id"`
		SubjectID int    `json:"subject_id"`
		StartsAt  string `json:"starts_at"`
		EndsAt    string `json:"ends_at"`
		Room      string `json:"room"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if input.GroupID <= 0 ||
		input.TeacherID <= 0 ||
		input.SubjectID <= 0 ||
		strings.TrimSpace(input.StartsAt) == "" ||
		strings.TrimSpace(input.EndsAt) == "" ||
		strings.TrimSpace(input.Room) == "" {
		writeError(w, http.StatusUnprocessableEntity, "invalid input")
		return
	}

	lesson, err := h.store.Create(
		r.Context(),
		input.GroupID,
		input.TeacherID,
		input.SubjectID,
		input.StartsAt,
		input.EndsAt,
		input.Room,
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusCreated, lesson)
}

func (h *LessonHandler) SetStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var input struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	input.Status = strings.TrimSpace(input.Status)

	if input.Status == "" {
		writeError(w, http.StatusUnprocessableEntity, "invalid status")
		return
	}

	err = h.store.SetStatus(r.Context(), id, input.Status)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "lesson not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	lesson, err := h.store.Get(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(w, http.StatusOK, lesson)
}

func (h *LessonHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.store.Delete(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "lesson not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}