package i_cant_beleave_it_can

import "constraints"

func Sort[T constraints.Ordered](v []T) {
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v); j++ {
			if v[i] < v[j] {
				v[i], v[j] = v[j], v[i]
			}
		}
	}
}
