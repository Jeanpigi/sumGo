package utils

// SumGo es la suma de valores enteros
func SumGo(valores ...int) int {
	total := 0
	for _, valor := range valores {
		total += valor
	}
	return total
}