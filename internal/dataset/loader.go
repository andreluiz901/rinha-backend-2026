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

	vectors := make([][14]float32, 0, estimated)
	labels := make([]int, 0, estimated)

	count := 0

	for decoder.More() {
		var item rawItem

		if err := decoder.Decode(&item); err != nil {
			return nil, err
		}

		// add vector (flatten)
		vectors = append(vectors, item.Vector)

		// label → uint8
		if item.Label == "fraud" {
			labels = append(labels, 1)
		} else {
			labels = append(labels, 0)
		}

		count++

		//  optional debug
		if count%500_000 == 0 {
			fmt.Println("Loaded:", count)
		}
	}

	fmt.Println("Total carregado:", count)

	return &Dataset{
		Vectors: vectors,
		Labels:  labels,
		Size:    count,
	}, nil
}