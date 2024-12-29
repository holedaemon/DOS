package null

type Nullable[T any] *T

func Some[T any](v T) Nullable[T] {
	return &v
}

func PtrTo[T any](v T) *T {
	return &v
}
