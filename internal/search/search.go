package search

import (
    "rinha-fraude/internal/types"
    "sort"
)

func distance(a, b [14]float32) float32 {
    var sum float32
    for i := 0; i < 14; i++ {
        diff := a[i] - b[i]
        sum += diff * diff
    }
    return sum
}

func TopK(vectors []float32, labels []uint8, size int, query [14]float32, k int) []types.Neighbor {
    neighbors := make([]types.Neighbor, 0, size)

    for i := 0; i < size; i++ {
        start := i * 14

        var vec [14]float32
        copy(vec[:], vectors[start:start+14])

        dist := distance(query, vec)

        neighbors = append(neighbors, types.Neighbor{
            Distance: dist,
            Label:    labels[i],
        })
    }

    sort.Slice(neighbors, func(i, j int) bool {
        return neighbors[i].Distance < neighbors[j].Distance
    })

    return neighbors[:k]
}