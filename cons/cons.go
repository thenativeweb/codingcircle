package cons

type PairFunc func(int, int) int

func Cons(a, b int) func(PairFunc) int {
	return func(f PairFunc) int {
		return f(a, b)
	}
}

func Car(p func(PairFunc) int) int {
	return p(func(a, b int) int {
		return a
	})
}

func Cdr(p func(PairFunc) int) int {
	return p(func(a, b int) int {
		return b
	})
}
