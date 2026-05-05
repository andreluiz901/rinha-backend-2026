package dataset

type Dataset struct {
    Vectors []float32
    Labels  []uint8
    Size    int
}

// fake dataset to test
func NewMockDataset(n int) *Dataset {
    vectors := make([]float32, n*14)
    labels := make([]uint8, n)

    for i := 0; i < n; i++ {
        for j := 0; j < 14; j++ {
            vectors[i*14+j] = float32(i % 10) // simple patter
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