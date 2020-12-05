package array

// Sum takes an array of floating point numbers as input and adds them to return the sum
func Sum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// SumAll takes a slice of slice of floating point numbers and
// returns a slice containing the sum of individual slice elements
func SumAll(listOfNumberList [][]float64) (listSum []float64) {
	for _, list := range listOfNumberList {
		listSum = append(listSum, Sum(list))
	}
	return
}

// SumAllTails calculates sum of all elements of a slice except it's head element
func SumAllTails(numbersToSum ...[]float64) (sumList []float64) {
	for _, numbers := range numbersToSum {
		if len(numbers) <= 1 {
			sumList = append(sumList, 0)
		} else {
			sumList = append(sumList, Sum(numbers[1:]))
		}
	}
	return
}
