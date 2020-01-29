package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	if numbers != nil && len(numbers) > 0 {
		sum := 0
		for _, num := range numbers {
			sum += num
		}
		return sum
	}
	return 0
}
