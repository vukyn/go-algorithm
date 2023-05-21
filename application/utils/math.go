package utils

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(vals ...int) int {
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func MinF(vals ...float64) float64 {
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}