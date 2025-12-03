package seqs

import "iter"

func Map[T, O any](in iter.Seq[T], f func(T) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for t := range in {
			if !yield(f(t)) {
				break
			}
		}
	}
}
