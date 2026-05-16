package api

import (
	"encoding/json"
	"net/http"
	//"time"
	//"fmt"
	//"os"

	"rinha-fraude/internal/engine"
	"rinha-fraude/internal/types"
)

type Handler struct {
	Engine *engine.Engine
}

type FraudResponse struct {
	Approved   bool    `json:"approved"`
	FraudScore float32 `json:"fraud_score"`
}

func (h *Handler) FraudScore(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req types.FraudRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		//fmt.Println("decode error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	//start := time.Now()
	approved, fraudScore := h.Engine.Predict(req)
	//elapsed := time.Since(start)
	// const debug = true
	// if debug {
	// 	fmt.Println("predict:", elapsed)
	// }

	resp := FraudResponse{
		Approved:   approved,
		FraudScore: fraudScore,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func Ready(w http.ResponseWriter, r *http.Request) {

	if !ReadyState.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}