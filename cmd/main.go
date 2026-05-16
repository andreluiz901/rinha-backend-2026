package main

import (
    "fmt"
    "rinha-fraude/internal/dataset"
    "rinha-fraude/internal/vector"
    "rinha-fraude/internal/config"
    "rinha-fraude/internal/engine"
    "net/http"
    "rinha-fraude/internal/api"
)


func main() {

    normalization := config.DefaultNormalization()

    vector.SetMccRisk(config.DefaultMccRisk())

    //fmt.Println("Initing dataset...")
    api.ReadyState.Store(false)
    ds, err := dataset.LoadDataset("resources/references.json.gz")

    if err != nil {
        panic(err)
    }
    //fmt.Println("Dataset loaded!")
    api.ReadyState.Store(true)

    eng := &engine.Engine{
        Dataset: ds,
        Normalization: normalization,
    }

    handler := &api.Handler{
		Engine: eng,
	}

    http.HandleFunc("/fraud-score", handler.FraudScore)
    http.HandleFunc("/ready", api.Ready)

    fmt.Println("Server listening on :9999")

    err = http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}

}