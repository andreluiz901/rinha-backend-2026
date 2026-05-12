package dataset

type Dataset struct {
    Vectors [][14]float32
    Labels  []int
    Size    int

    CoarseIndex  map[string][]int   // specific bucket
	BroadIndex   map[string][]int   // broad bucket - fallback
}

// fake dataset to test
func NewMockDataset(n int) *Dataset {
    vectors := make([][14]float32, n)
    labels := make([]int, n)

    for i := 0; i < n; i++ {
        for j := 0; j < 14; j++ {
            vectors[i][j] = float32(i % 10) // padrão simples
        }

        if i%2 == 0 {
            labels[i] = 1 // fraud
        } else {
            labels[i] = 0
        }
    }

    return &Dataset{
        Vectors: vectors,
        Labels:  labels,
        Size:    n,
    }
}