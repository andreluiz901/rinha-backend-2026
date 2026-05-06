package score

import "rinha-fraude/internal/search"

func FraudScore(neighbors []search.Neighbor) float32 {
    frauds := 0

    for _, n := range neighbors {
        if n.Label == 1 {
            frauds++
        }
    }

    return float32(frauds) / float32(len(neighbors))
}