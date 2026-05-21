package engine

import (
	"rinha-fraude/internal/dataset"
	"rinha-fraude/internal/score"
	"rinha-fraude/internal/search"
	"rinha-fraude/internal/types"
	"rinha-fraude/internal/vector"
	//"fmt"
	//"time"
)

type Engine struct {
	Dataset      *dataset.Dataset
	Normalization vector.Normalization
}

func (e *Engine) Predict(req types.FraudRequest) (bool, float32) {
	//start1 := time.Now()
	vec := vector.BuildVector(req, e.Normalization)
	//fmt.Println("vector:", time.Since(start1))

	candidates := e.Dataset.Candidates(vec)

	//fmt.Println("candidates:", len(candidates))
	//start2 := time.Now()
	neighbors := search.TopK(
		e.Dataset.Vectors,
		e.Dataset.Labels,
		vec,
		5,
		candidates,
	)
	//fmt.Println("topk:", time.Since(start2))

	fraudScore := score.FraudScore(neighbors)

	approved := fraudScore < 0.6

	return approved, fraudScore
}