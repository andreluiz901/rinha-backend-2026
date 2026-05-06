package main

import (
    "fmt"
    // "rinha-fraude/internal/dataset"
    // "rinha-fraude/internal/search"
    // "rinha-fraude/internal/score"
    "rinha-fraude/internal/types"
    "rinha-fraude/internal/vector"
)

func buildTestInput() types.FraudRequest {
    return types.FraudRequest{
        Transaction: types.Transaction{
            Amount:       41.12,
            Installments: 2,
            RequestedAt:  "2026-03-11T18:45:53Z",
        },
        Customer: types.Customer{
            AvgAmount:      82.24,
            TxCount24h:     3,
            KnownMerchants: []string{"MERC-001"},
        },
        Merchant: types.Merchant{
            ID:        "MERC-001",
            MCC:       "5912",
            AvgAmount: 298.95,
        },
        Terminal: types.Terminal{
            IsOnline:    false,
            CardPresent: true,
            KmFromHome:  29.23,
        },
        LastTransaction: nil,
        // LastTransaction: &types.LastTransaction{
        //     Timestamp:     "2026-03-11T17:45:53Z",
        //     KmFromCurrent: 10,
        // },
    }
}

func loadMccRisk() map[string]float32 {
	return map[string]float32{
		"5411": 0.15,
		"5812": 0.30,
		"5912": 0.20,
		"5944": 0.45,
		"7801": 0.80,
		"7802": 0.75,
		"7995": 0.85,
		"4511": 0.35,
		"5311": 0.25,
		"5999": 0.50,
	}
}

func main() {
    // fmt.Println("Initing dataset...")

    // ds, err := dataset.LoadDataset("resources/references.json.gz")

    // if err != nil {
    //     panic(err)
    // }

    // fmt.Println("Dataset loaded!")

    // input := types.TransactionInput{
    //     Amount: 5,
    // }

    // vec := vector.BuildVector(input)

    // neighbors := search.TopK(ds.Vectors, ds.Labels, ds.Size, vec, 5)

    // fraudScore := score.FraudScore(neighbors)

    // approved := fraudScore < 0.6

    // fmt.Println("Result:")
    // fmt.Println("Score:", fraudScore)
    // fmt.Println("Approved:", approved)

    // print test vector/normalized vector

    normalization := vector.Normalization{
        MaxAmount: 10000,
        MaxInstallments:  12,
        AmountVsAvgRatio: 10,
        MaxKm: 1000,
        MaxMinutes: 1440,
        MaxTxCount24h: 20,
        MaxMerchantAvgAmount:   10000,
    }
    
    input := buildTestInput()
    
    vector.SetMccRisk(loadMccRisk())

    vec := vector.BuildVector(input, normalization)
    
    fmt.Println("Vector:", vec)
    // fmt.Printf("Vector: %+v\n", vec)
    // for i, v := range vec {
    //    fmt.Printf("[%d] = %.6f\n", i, v)
    // }
    
}