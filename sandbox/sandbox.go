package sandbox

type T int

func (t *T) Mutate(value T) {
	(*t) = value
}
