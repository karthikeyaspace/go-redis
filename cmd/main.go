package main

import (
	"log"
	"net/http"

	"github.com/karthikeyaspace/game-leaderboard/internal/config"
	"github.com/karthikeyaspace/game-leaderboard/internal/db"
	"github.com/karthikeyaspace/game-leaderboard/internal/handlers"
	"github.com/karthikeyaspace/game-leaderboard/internal/services"
	"github.com/karthikeyaspace/game-leaderboard/internal/utils"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

var cfg = config.NewConfig()

func (s *APIServer) Start() error {

	db, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	defer db.Close()

	service := services.NewService(db)
	handler := handlers.NewHandler(service)

	router := http.NewServeMux()

	router.HandleFunc("GET /", handler.HomeHandler)
	router.HandleFunc("POST /create-player", handler.CreatePlayerHandler)
	router.HandleFunc("POST /update-score", handler.UpdateScoreHandler)
	router.HandleFunc("GET /leaderboard", handler.GetLeaderboardHandler)

	middleware := utils.Logger(utils.CORS(router))

	server := &http.Server{
		Addr:    s.addr,
		Handler: middleware,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}

	return nil
}

func main() {
	server := NewAPIServer(cfg.Port)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Printf("Starting API server at http://localhost:%s", cfg.Port)
}
