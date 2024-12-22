package handlers

import (
	"net/http"

	"github.com/karthikeyaspace/game-leaderboard/internal/services"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) UpdateScoreHandler(w http.ResponseWriter, r *http.Request) {
	// h.service.UpdateScore("someUserId", 100)
	w.Write([]byte("Score updated"))
}

func (h *Handler) GetLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	// h.service.GetLeaderboard()
	w.Write([]byte("Leaderboard"))
}

func (h *Handler) HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Home"))
}
