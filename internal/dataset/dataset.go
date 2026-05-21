package dataset

type Dataset struct {
    Vectors []Vector
    Labels  []uint8
    Size    int

    CoarseIndex  map[uint32][]uint32   // specific bucket
	BroadIndex   map[uint32][]uint32   // broad bucket - fallback
}

// fake dataset to test
// func NewMockDataset(n int) *Dataset {
//     vectors := make([][14]float32, n)
//     labels := make([]int, n)

//     for i := 0; i < n; i++ {
//         for j := 0; j < 14; j++ {
//             vectors[i][j] = float32(i % 10) // padrão simples
//         }

//         if i%2 == 0 {
//             labels[i] = 1 // fraud
//         } else {
//             labels[i] = 0
//         }
//     }

//     return &Dataset{
//         Vectors: vectors,
//         Labels:  labels,
//         Size:    n,
//     }
// }