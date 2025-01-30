package errors_examples

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
