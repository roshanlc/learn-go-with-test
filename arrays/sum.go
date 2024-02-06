package arrays

// Sum returns the sum of numbers in the given array
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}

// SumAll returns the sum of each slices
func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, val := range numbers {
		sums = append(sums, Sum(val))
	}
	return sums
}
