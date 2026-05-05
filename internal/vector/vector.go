package vector

import "rinha-fraude/internal/types"

type Normalization struct {
	MaxAmount        float32
	MaxInstallments  float32
	AmountVsAvgRatio float32
}

func clamp(v float32) float32 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

// TEMP — placeholder to be replaced with the real logic of the 14 dimensions
func BuildVector(input types.FraudRequest, norm Normalization) [14]float32 {
    var v [14]float32

    // fake  example to test
    v[0] = input.Amount

    return v
}