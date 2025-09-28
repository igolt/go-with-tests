package generic_arrays_and_slices

func Reduce[T any, V any](
	collection []T,
	reducer func(accumulator V, currentValue T, index int) V,
	accumulator V,
) V {
	for index, value := range collection {
		accumulator = reducer(accumulator, value, index)
	}
	return accumulator
}
