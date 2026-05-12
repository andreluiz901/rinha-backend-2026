package engine

import (
	"rinha-fraude/internal/dataset"
	"rinha-fraude/internal/score"
	"rinha-fraude/internal/search"
	"rinha-fraude/internal/types"
	"rinha-fraude/internal/vector"
	//"fmt"
)

type Engine struct {
	Dataset      *dataset.Dataset
	Normalization vector.Normalization
}

func (e *Engine) Predict(req types.FraudRequest) (bool, float32) {
	vec := vector.BuildVector(req, e.Normalization)

	candidates := e.Dataset.Candidates(vec)

	//fmt.Println("candidates:", len(candidates))
	
	neighbors := search.TopK(
		e.Dataset.Vectors,
		e.Dataset.Labels,
		vec,
		5,
		candidates,
	)

	fraudScore := score.FraudScore(neighbors)

	approved := fraudScore < 0.6

	return approved, fraudScore
}