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

func Rotate[T any](s []T, middle int) []T {
	s = append(s[middle:], s[:middle]...)
	return s
}

func CountFunc[T any](s []T, f func(value T) bool) int {
	return Reduce(s, 0, func(acc int, v T) int {
		if f(v) {
			return acc + 1
		}
		return acc
	})
}

func Count[T comparable](s []T, value T) int {
	return CountFunc(s, func(v T) bool {
		return v == value
	})
}
