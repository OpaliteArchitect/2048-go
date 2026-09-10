package main

import (
	"encoding/json"
	"net/http"

	"github.com/opalitearchitect/2048-go/game"
)

type MoveRequest struct {
	Direction game.Direction `json:"direction"`
}

type GameResponse struct {
	Board game.Board `json:"board"`
	Turns int        `json:"turns"`
	Moved bool       `json:"moved"`
}

type Server struct {
	game game.Game
}

func (s *Server) HandleMove(w http.ResponseWriter, r *http.Request) {
	var req MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	moved := s.game.Move(req.Direction)

	resp := GameResponse{
		Board: s.game.GetBoard(),
		Turns: s.game.GetTurns(),
		Moved: moved,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "# 2048 Game API Server\n\nWelcome! Use this server to control the running 2048 game session over HTTP.", http.StatusOK)
}

func main() {
	g := game.New()
	server := Server{game: &g}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /move", server.HandleMove)
	mux.HandleFunc("GET /", server.HandleIndex)

	http.ListenAndServe(":2048", mux)
}
