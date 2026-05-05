package main

import (
    "fmt"
    "rinha-fraude/internal/dataset"
    "rinha-fraude/internal/search"
    "rinha-fraude/internal/score"
    "rinha-fraude/internal/types"
    "rinha-fraude/internal/vector"
)

func main() {
    fmt.Println("Initing dataset...")

    ds, err := dataset.LoadDataset("resources/references.json.gz")

    if err != nil {
        panic(err)
    }

    fmt.Println("Dataset loaded!")

    input := types.TransactionInput{
        Amount: 5,
    }

    vec := vector.BuildVector(input)

    neighbors := search.TopK(ds.Vectors, ds.Labels, ds.Size, vec, 5)

    fraudScore := score.FraudScore(neighbors)

    approved := fraudScore < 0.6

    fmt.Println("Result:")
    fmt.Println("Score:", fraudScore)
    fmt.Println("Approved:", approved)
}