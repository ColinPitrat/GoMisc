package factorial

func Factorial(n int) (result int) {
	result = 1
	for n > 1 {
		result *= n
		n--
	}
	return result
}
