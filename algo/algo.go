package algo

func Map[T any, O any](s []T, f func(T) O) []O {
	result := make([]O, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

func Reduce[T any, O any](s []T, init O, f func(acc O, v T) O) O {
	result := init
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

func Filter[T any](s []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
