package main

import (
    "fmt"
    "net/http"
    "time"

    "rinha-fraude/internal/dataset"
    "rinha-fraude/internal/vector"
    "rinha-fraude/internal/config"
    "rinha-fraude/internal/engine"
    "rinha-fraude/internal/api"

    _ "net/http/pprof"
    "runtime/debug"
)

var eng *engine.Engine

func main() {

    normalization := config.DefaultNormalization()

    vector.SetMccRisk(config.DefaultMccRisk())

    //fmt.Println("Initing dataset...")
    api.ReadyState.Store(false)

    handler := &api.Handler{}

    http.HandleFunc("/fraud-score", handler.FraudScore)
    http.HandleFunc("/ready", api.Ready)

    go func() {
        fmt.Println("Server listening on :9999")

        srv := &http.Server{
            Addr:              ":9999",
            Handler:           nil,
            ReadHeaderTimeout: 2 * time.Second,
            IdleTimeout:       30 * time.Second,
        }

        err := srv.ListenAndServe()
        if err != nil {
            panic(err)
        }
    }()

    go func() {
        http.ListenAndServe("0.0.0.0:6060", nil)
    }()

    ds, err := dataset.LoadDataset("resources/references.json.gz")
    if err != nil {
        panic(err)
    }
    //fmt.Println("Dataset loaded!")

    debug.FreeOSMemory()
    
    eng := &engine.Engine{
        Dataset: ds,
        Normalization: normalization,
    }
    
    handler.Engine = eng
    
    api.ReadyState.Store(true)
    
    select {}   

}