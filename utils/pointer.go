package utils

func ToPtr[T any](v T) *T {
	return &v
}

func FromPtr[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

func SafePtr[T any](p *T, defaultValue T) T {
	if p == nil {
		return defaultValue
	}
	return *p
}

func IsNilPtr[T any](p *T) bool {
	return p == nil
}

func NilPtr[T any]() *T {
	return nil
}
