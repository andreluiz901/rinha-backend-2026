package search

import (
    //"rinha-fraude/internal/types"
    //"sort"
)

/* 
    each vector, calculate distance, maitain 5 least

*/


type Neighbor struct {
    Distance float32
    Label    int // 1 = fraud, 0 = legit
}

// distance non euclidian, no square root to gain performance, test

func distance(a, b [14]float32) float32 {
    var sum float32
    for i := 0; i < 14; i++ {
        d := a[i] - b[i]
        sum += d * d
    }
    return sum
}

// simplified top k neigh, test

func TopK(vectors [][14]float32, labels []int, query [14]float32, k int) []Neighbor {
    neighbors := make([]Neighbor, 0, k)

    for i, v := range vectors {
        dist := distance(query, v)

        if len(neighbors) < k {
            neighbors = append(neighbors, Neighbor{dist, labels[i]})
            continue
        }

        // find worse (major distance)
        maxIdx := 0
        for j := 1; j < k; j++ {
            if neighbors[j].Distance > neighbors[maxIdx].Distance {
                maxIdx = j
            }
        }

        if dist < neighbors[maxIdx].Distance {
            neighbors[maxIdx] = Neighbor{dist, labels[i]}
        }
    }

    return neighbors
}