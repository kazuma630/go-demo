package mylib

func Average(s []int) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return int(sum / len(s))
}
