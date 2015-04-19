package goj

func Each(fn func(value interface{}, key int), args ...[]interface{}) {
	for _, arr := range args {
		for k, v := range arr {
			fn(v, k)
		}
	}
}

func Map(f Function, input []T) []T {
	result := make([]T, len(input))

	for i, element := range input {
		result[i] = f(element)
	}

	return result
}

func Filter(fn FilterFn, input []T) (result []T) {
	for _, v := range input {
		if fn(v) {
			result = append(result, v)
		}
	}
	return
}
