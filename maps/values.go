package maps

func ValuesFiltered[K comparable, V any](m map[K]V, fn func(V) bool) []V {
	var vs []V

	for _, v := range m {
		if fn(v) {
			vs = append(vs, v)
		}
	}

	return vs
}

func Values[K comparable, V any](m map[K]V) []V {
	return ValuesFiltered(m, func(k V) bool {
		return true
	})
}
