package search

import "rinha-fraude/internal/dataset"

type Neighbor struct {
	Distance float32
	Label    uint8
}

// distance non euclidian, no square root to gain performance, test

func distance(a [14]float32, b dataset.Vector) float32 {
	var sum float32

	
	for i := 0; i < 14; i++ {
		bv := dataset.Dequantize(b[i])
		d := a[i] - bv
		sum += d * d
	}

	return sum
}

// simplified top k neigh, test

func TopK(
	vectors []dataset.Vector,
	labels []uint8,
	query [14]float32,
	k int,
  candidates []int,
) []Neighbor {

	neighbors := make([]Neighbor, 0, k)

	// fallback → scan total
	if len(candidates) == 0 {

		for i := 0; i < len(vectors); i++ {

			dist := distance(query, vectors[i])

			if len(neighbors) < k {

				neighbors = append(neighbors, Neighbor{
					Distance: dist,
					Label: labels[i],
				})

				continue
			}

			worst := 0

			for j := 1; j < k; j++ {
				if neighbors[j].Distance > neighbors[worst].Distance {
					worst = j
				}
			}

			if dist < neighbors[worst].Distance {
				neighbors[worst] = Neighbor{
					Distance: dist,
					Label: labels[i],
				}
			}
		}

		return neighbors
	}

	// candidates-only scan
	for _, idx := range candidates {

		dist := distance(query, vectors[idx])

		if len(neighbors) < k {

			neighbors = append(neighbors, Neighbor{
				Distance: dist,
				Label: labels[idx],
			})

			continue
		}

		worst := 0

		for j := 1; j < k; j++ {
			if neighbors[j].Distance > neighbors[worst].Distance {
				worst = j
			}
		}

		if dist < neighbors[worst].Distance {

			neighbors[worst] = Neighbor{
				Distance: dist,
				Label: labels[idx],
			}
		}
	}

	return neighbors   
}