package dataset

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
)

type rawItem struct {
	Vector [14]float32 `json:"vector"`
	Label  string      `json:"label"`
}

func LoadDataset(path string) (*Dataset, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	decoder := json.NewDecoder(gz)

	// read the start of the array '['
	t, err := decoder.Token()
	if err != nil {
		return nil, err
	}

	if delim, ok := t.(json.Delim); !ok || delim != '[' {
		return nil, fmt.Errorf("expected JSON array")
	}

	// initial estimative (avoid reallocation)
	const estimated = 3_000_000

	vectors := make([]Vector, 0, estimated)
	labels := make([]uint8, 0, estimated)
	
	coarseIndex := make(map[uint32][]uint32)
	broadIndex := make(map[uint32][]uint32)

	count := 0

	for decoder.More() {
		var item rawItem

		if err := decoder.Decode(&item); err != nil {
			return nil, err
		}

		
		var compact Vector

		for i := 0; i < 14; i++ {
			compact[i] = Quantize(item.Vector[i])
		}

		vectors = append(vectors, compact)

		idx := uint32(count)

		// label → uint8
		if item.Label == "fraud" {
			labels = append(labels, 1)
		} else {
			labels = append(labels, 0)
		}

		ck := coarseKey(item.Vector)
		bk := broadKey(item.Vector)

		coarseIndex[ck] = append(coarseIndex[ck], idx)
		broadIndex[bk] = append(broadIndex[bk], idx)

		count++

		//  optional debug
		if count%500_000 == 0 {
			fmt.Println("Loaded:", count)
		}
	}

	fmt.Println("Total carregado:", count)

	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)

	// fmt.Printf("HeapAlloc = %.2f MB\n", float64(m.HeapAlloc)/1024/1024)
	// fmt.Printf("Sys = %.2f MB\n", float64(m.Sys)/1024/1024)
	// fmt.Printf("NumGC = %d\n", m.NumGC)
	return &Dataset{
		Vectors: vectors,
		Labels:  labels,
		Size:    count,

		CoarseIndex: coarseIndex,
		BroadIndex: broadIndex,
	}, nil
}