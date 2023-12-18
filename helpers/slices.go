package helpers

func Map[T, U any](input []T, mapper func(T) U) []U {
	u := make([]U, 0)
	for _, i := range input {
		u = append(u, mapper(i))
	}
	return u
}

func GroupBy[T any, K comparable](input []T, grouper func(T) K) map[K][]T {
	m := make(map[K][]T)
	for _, i := range input {
		key := grouper(i)
		_, isExist := m[key]
		if !isExist {
			m[key] = []T{i}
		} else {
			m[key] = append(m[key], i)
		}
	}
	return m
}
