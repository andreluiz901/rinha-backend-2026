package score

import "rinha-fraude/internal/types"

func FraudScore(neighbors []types.Neighbor) float32 {
    var frauds int

    for _, n := range neighbors {
        if n.Label == 1 {
            frauds++
        }
    }

    return float32(frauds) / float32(len(neighbors))
}