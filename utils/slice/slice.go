package slice

func Map[T1, T2 interface{}](list []T1, mapFunc func(int, T1) T2) []T2 {
	newList := make([]T2, 0, len(list))
	for i, v := range list {
		newList = append(newList, mapFunc(i, v))
	}
	return newList
}

func Filter[T interface{}](list []T, filterFunc func(int, T) bool) []T {
	newList := make([]T, 0, len(list))
	for i, v := range list {
		if filterFunc(i, v) {
			newList = append(newList, v)
		}
	}
	return newList
}
