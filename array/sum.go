package array

// Sum takes an array of floating point numbers as input and adds them to return the sum
func Sum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
