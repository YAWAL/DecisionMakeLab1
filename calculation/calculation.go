package calculation

// Необхідно знайти оптимальні стратегії при песимістичній оцінці (по критерію
// 	Вальда), оцінці Лапласа, по критерію Гурвіца. Значення коефіцієнта оптимізму
// 	вибрати самостійно. Результати вибору рішення відобразити в таблиці. Зробити
// 	висновки по застосуванню критеріїв

const (
	optCoef      float32 = 0.3
	possibility1 float32 = 0.5
	possibility2 float32 = 0.35
	possibility3 float32 = 0.15
)

var possibilities = map[int]float32{
	0: possibility1,
	1: possibility2,
	2: possibility3,
}

func Gurvic(input [3][3]int) map[int]float32 {
	gurvic := make(map[int]float32)
	for i, alts := range input {
		min := min(alts)
		max := max(alts)
		gurvicValue := optCoef*min + (1-optCoef)*max
		gurvic[i] = gurvicValue
	}
	return gurvic
}

func Wald(input [3][3]int) map[int]float32 {
	wald := make(map[int]float32)
	for i, alts := range input {
		wald[i] = min(alts)
	}
	return wald
}

func Laplas(input [3][3]int) map[int]float32 {
	laplas := make(map[int]float32)
	for i, alts := range input {
		laplas[i] = sum(alts)
	}
	return laplas
}

func BayerLaplas(input [3][3]int) map[int]float32 {
	bayerLaplas := make(map[int]float32)
	for i, alts := range input {
		bayerLaplas[i] = brLpls(alts, possibilities[i])
	}
	return bayerLaplas
}

func brLpls(values [3]int, possibility float32) float32 {
	var sum float32
	for _, elem := range values {
		sum += float32(elem) * possibility
	}
	return sum
}

func max(values [3]int) float32 {
	max := values[0]
	for _, elem := range values {
		if elem > max {
			max = elem
		}
	}
	return float32(max)
}

func min(values [3]int) float32 {
	min := values[0]
	for _, elem := range values {
		if elem < min {
			min = elem
		}
	}
	return float32(min)
}

func sum(values [3]int) float32 {
	var sum float32
	for _, elem := range values {
		sum = sum + float32(elem)
	}
	return sum
}
