package variadic_reduction

func KurangVariadic(numbers ...int) int {
	result := numbers[0]
	for _, n := range numbers[1:] {
		result -= n
	}
	return result
}