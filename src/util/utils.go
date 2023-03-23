package util

func factorial(i int) (y int) {
	if i < 0 {
		return -1.
	}
	y = 1
	for u := 2; u <= i; u++ {
		y = u * y
	}
	return y
}
