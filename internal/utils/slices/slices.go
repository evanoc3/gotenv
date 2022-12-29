package slices


func Transform[K any, V any](items []K, transformFunc func(K) V) []V {
	transformed := make([]V, len(items))

	for i, item := range items {
		transformed[i] = transformFunc(item)
	}

	return transformed
}


func Filter[T comparable](items []T, filterFunc func(T) bool) []T {
	indexesOfFilteredItems := make([]int, 0, len(items))

	for i, item := range items {
		if filterFunc(item) {
			indexesOfFilteredItems = append(indexesOfFilteredItems, i)
		}
	}

	filteredItems := make([]T, 0, len(indexesOfFilteredItems))
	for _, i := range(indexesOfFilteredItems) {
		filteredItems = append(filteredItems, items[i])
	}

	return filteredItems
}
