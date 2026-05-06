package vector

import (
    "rinha-fraude/internal/types"
    "time"
)

type Normalization struct {
	MaxAmount        float32
	MaxInstallments  float32
	AmountVsAvgRatio float32
    MaxMinutes             float32
	MaxKm                  float32
	MaxTxCount24h          float32
	MaxMerchantAvgAmount   float32
}

func clamp(v float32) float32 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

var mccRiskMap map[string]float32

func SetMccRisk(m map[string]float32) {
    mccRiskMap = m
}

// TEMP — placeholder to be replaced with the real logic of the 14 dimensions
func BuildVector(input types.FraudRequest, norm Normalization) [14]float32 {
   
    // normalize time
    var v [14]float32

    t, err := time.Parse(time.RFC3339, input.Transaction.RequestedAt) // converts string in real time
    if err != nil {
        return v // fallback just in case, return empty vector
        // TODO: error treatment
    }

   // [0] amount
	v[0] = clamp(input.Transaction.Amount / norm.MaxAmount)

	// [1] installments
	v[1] = clamp(float32(input.Transaction.Installments) / norm.MaxInstallments)

	// [2] amount_vs_avg
	if input.Customer.AvgAmount > 0 {
		ratio := input.Transaction.Amount / input.Customer.AvgAmount
		v[2] = clamp(ratio / norm.AmountVsAvgRatio)
	} else {
		v[2] = 0
	}

    // [3] hour_of_day
    hour := t.Hour()

    v[3] = float32(hour) / 23

    // [4] day_of_week --> with some corrections
    weekday := int(t.Weekday())

    // convert go date patter into the one used at dataset
    adjusted := (weekday + 6) % 7

    v[4] = float32(adjusted) / 6

    // [5] minutes_since_last_tx
    if input.LastTransaction == nil {
        v[5] = -1
    } else {
        lastTime, err := time.Parse(time.RFC3339, input.LastTransaction.Timestamp)
        if err != nil {
            v[5] = -1
        } else {
            minutes := t.Sub(lastTime).Minutes()

            if minutes < 0 {
                minutes = 0
            }

            v[5] = clamp(float32(minutes) / norm.MaxMinutes)
        }
    }

    // [6] km_from_last_tx
    if input.LastTransaction == nil {
        v[6] = -1
    } else {
        v[6] = clamp(input.LastTransaction.KmFromCurrent / norm.MaxKm)
    }


    // [7] km_from_home
    v[7] = clamp(input.Terminal.KmFromHome / norm.MaxKm)

    // [8] tx_count_24h
    txCount := float32(input.Customer.TxCount24h)
    v[8] = clamp(txCount / norm.MaxTxCount24h)

    // [9] is_online
    if input.Terminal.IsOnline {
        v[9] = 1
    } else {
        v[9] = 0
    }

    // [10] card_present
    if input.Terminal.CardPresent {
        v[10] = 1
    } else {
        v[10] = 0
    }

    // [11] unknown_merchant
    // 1 → merchant not in known_merchants
    // 0 → merchant is in known_merchant list

    found := false
    for _, m := range input.Customer.KnownMerchants {
        if m == input.Merchant.ID {
            found = true
            break
        }
    }

    if found {
        v[11] = 0
    } else {
        v[11] = 1
    }

    // [12] mcc_risk
    // MCC map
    risk := float32(0.5)

    if mccRiskMap != nil {
        if r, ok := mccRiskMap[input.Merchant.MCC]; ok {
            risk = r
        }
    }

    v[12] = risk

    // [13] merchant_avg_amount

    v[13] = clamp(input.Merchant.AvgAmount / norm.MaxMerchantAvgAmount)
    


	return v
}