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

var approvedResponse = []byte(`{"approved":true,"fraud_score":1}`)
var deniedResponse = []byte(`{"approved":false,"fraud_score":0}`)

type Handler struct {
	Engine *engine.Engine
}

type FraudResponse struct {
	Approved   bool    `json:"approved"`
	FraudScore float32 `json:"fraud_score"`
}

func (h *Handler) FraudScore(w http.ResponseWriter, r *http.Request) {

	if !ReadyState.Load() || h.Engine == nil {
		http.Error(w, "starting", http.StatusServiceUnavailable)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// var req types.FraudRequest
	req := types.FraudRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// err := json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	//fmt.Println("decode error:", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	
	//start := time.Now()
	approved, _ := h.Engine.Predict(req)
	//approved := true
	//elapsed := time.Since(start)
	// const debug = true
	// if debug {
	// 	fmt.Println("predict:", elapsed)
	// }

	// resp := FraudResponse{
	// 	Approved:   approved,
	// 	FraudScore: fraudScore,
	// }

	w.Header().Set("Content-Type", "application/json")

	//json.NewEncoder(w).Encode(resp)

	if approved {
		w.Write(approvedResponse)
	} else {
		w.Write(deniedResponse)
	}
}

func Ready(w http.ResponseWriter, r *http.Request) {

	if !ReadyState.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}