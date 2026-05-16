package dataset

type Vector [14]uint8

func Quantize(v float32) uint8 {
	if v < 0 {
		return 0
	}

	if v > 1 {
		v = 1
	}

	return uint8(v * 255)
}

func Dequantize(v uint8) float32 {
	return float32(v) / 255.0
}

var DequantTable [256]float32

func init() {
	for i := 0; i < 256; i++ {
		DequantTable[i] = float32(i) / 255.0
	}
}