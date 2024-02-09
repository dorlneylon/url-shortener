package utils

import "errors"

type Either[L, R any] struct {
	Left  L
	Right R
}

func NewEither[L, R any](left L, right R) *Either[L, R] {
	return &Either[L, R]{
		Left:  left,
		Right: right,
	}
}

func (e *Either[L, R]) IsLeft() bool {
	return e.Left != nil
}

func (e *Either[L, R]) IsRight() bool {
	return e.Right != nil
}

func (e *Either[L, R]) FlatMap(f func(L) (*Either[L, R], error)) (*Either[L, R], error) {
	if e.IsLeft() {
		return f(e.Left)
	}
	return e, errors.New("Either is right")
}
