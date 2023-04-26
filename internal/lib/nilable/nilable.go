package nilable

func Map[T any, U any](t *T, f func(T) U) *U {
	if t == nil {
		return nil
	}
	u := f(*t)
	return &u
}
