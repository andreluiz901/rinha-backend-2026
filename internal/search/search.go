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
		bv := dataset.DequantTable[b[i]] // change strategy to avoid tail data(?)
		d := a[i] - bv
		sum += d * d
	}

	return sum
}

func TopK(
	vectors []dataset.Vector,
	labels []uint8,
	query [14]float32,
	k int,
  candidates []uint32,
) []Neighbor {

	// no candidates → immediate return(abort)
	if len(candidates) == 0 {
		return nil
	}

	//neighbors := make([]Neighbor, 0, k)
	var neighbors [5]Neighbor
	count := 0

	// candidates-only scan
	for _, idx := range candidates {
		i := int(idx)

		dist := distance(query, vectors[i])

		// filling topK
		if count < k {

			neighbors[count] = Neighbor{
				Distance: dist,
				Label:    labels[i],
			}

			count++
			continue
		}

		// find worst neigh
		worst := 0

		for j := 1; j < k; j++ {
			if neighbors[j].Distance > neighbors[worst].Distance {
				worst = j
			}
		}

		// substitutes if best
		if dist < neighbors[worst].Distance {

			neighbors[worst] = Neighbor{
				Distance: dist,
				Label: labels[i],
			}
		}
	}

	return neighbors[:count]
}