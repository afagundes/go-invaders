package utils

func ValueMinusPercent(val int, percentage float64) int {
	return int(float64(val) - float64(val)*percentage)
}
