package generic_arrays_and_slices

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	reducer := func(accumulator float64, currentValue Transaction, _ int) float64 {
		if currentValue.From == name {
			accumulator -= currentValue.Sum
		} else if currentValue.To == name {
			accumulator += currentValue.Sum
		}
		return accumulator
	}

	return Reduce(transactions, reducer, 0)
}
