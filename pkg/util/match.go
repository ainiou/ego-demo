package utils

func RemoveSlice[S comparable](all []S, remove []S) []S {
	removeMap := make(map[S]struct{})
	for _, v := range remove {
		removeMap[v] = struct{}{}
	}

	// 遍历 all slice，并将不在 removeMap 中的元素加入新的结果 slice 中
	var result []S
	for _, v := range all {
		if _, ok := removeMap[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

func SliceToMap[K comparable, S any](slice []S, mapFunc func(S) K) map[K]S {
	m := make(map[K]S)
	for _, s := range slice {
		m[mapFunc(s)] = s
	}
	return m
}
