package generic_arrays_and_slices

func Sum(numbers []int) int {
	return Reduce(numbers, sumReducer, 0)
}

func sumReducer(accumulator, currentValue, _ int) int {
	return accumulator + currentValue
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	return Reduce(numbersToSum, sumAllReducer, sums)
}

func sumAllReducer(accumulator, currentValue []int, index int) []int {
	accumulator[index] = Sum(currentValue)
	return accumulator
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	return Reduce(numbersToSum, sumAllTailsReducer, sums)
}

func sumAllTailsReducer(accumulator []int, currentValue []int, index int) []int {
	if len(currentValue) > 0 {
		accumulator[index] = Sum(currentValue[1:])
	}
	return accumulator
}
