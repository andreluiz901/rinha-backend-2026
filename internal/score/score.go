package score

import "rinha-fraude/internal/search"

func FraudScore(neighbors []search.Neighbor) float32 {
    frauds := 0

    for _, n := range neighbors {
        if n.Label == 1 {
            frauds++
        }
    }

    if len(neighbors) == 0 {
        return 1
    }
    
    return float32(frauds) / float32(len(neighbors))
}