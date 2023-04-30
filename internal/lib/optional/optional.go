package optional

func Map[T any, U any](t *T, f func(T) U) *U {
	if t == nil {
		return nil
	}
	u := f(*t)
	return &u
}

func FlatMap[T any, U any](t *T, f func(T) *U) *U {
	if t == nil {
		return nil
	}
	return f(*t)
}

func Filter[T any](t *T, f func(T) bool) *T {
	if t == nil {
		return nil
	}
	if !f(*t) {
		return nil
	}
	return t
}

func Reduce[T any, U any](t *T, onSome func(T) U, onNone func() U) U {
	if t == nil {
		return onNone()
	}
	return onSome(*t)
}
