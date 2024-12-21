package combinatorics

import "iter"

func Permutations[T any](s []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(s) == 0 {
			if !yield(s) {
				return
			}
		}
		// Steinhaus, implemented with a recursive closure.
		// arg is number of positions left to permute.
		// pass in len(s) to start generation.
		// on each call, weave element at pp through the elements 0..np-2,
		// then restore array to the way it was.
		var rc func(int)
		rc = func(np int) {
			if np == 1 {
				if !yield(s) {
					return
				}
			}
			np1 := np - 1
			pp := len(s) - np1
			// weave
			rc(np1)
			for i := pp; i > 0; i-- {
				s[i], s[i-1] = s[i-1], s[i]
				rc(np1)
			}
			// restore
			w := s[0]
			copy(s, s[1:pp+1])
			s[pp] = w
		}
		rc(len(s))
	}
}
