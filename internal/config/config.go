package config

import "rinha-fraude/internal/vector"

func DefaultNormalization() vector.Normalization {
	return vector.Normalization{
		MaxAmount:           10000,
		MaxInstallments:     12,
		AmountVsAvgRatio:    10,
		MaxKm:               1000,
		MaxMinutes:          1440,
		MaxTxCount24h:       20,
		MaxMerchantAvgAmount: 10000,
	}
}

func DefaultMccRisk() map[string]float32 {
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