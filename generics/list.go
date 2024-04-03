package generics

import "errors"

type List[T any] struct {
	data []T
}

func NewList[T any]() *List[T] {
	return &List[T]{
		data: make([]T, 0),
	}
}

func (l *List[T]) Append(value T) {
	l.data = append(l.data, value)
}

func (l *List[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(l.data) {
		var zero T
		return zero, errors.New("index out of range")
	}

	return l.data[index], nil
}

func (l *List[T]) Len() int {
	return len(l.data)
}

func (l *List[T]) Remove(index int) error {
	if index < 0 || index >= len(l.data) {
		return errors.New("index out of range")
	}
	l.data = append(l.data[:index], l.data[index+1:]...)
	return nil
}

func (l *List[T]) Set(index int, value T) error {
	if index < 0 || index >= len(l.data) {
		return errors.New("index out of range")
	}
	l.data[index] = value
	return nil
}

func (l *List[T]) Swap(i, j int) error {
	if i < 0 || i >= len(l.data) {
		return errors.New("index i out of range")
	}
	l.data[i], l.data[j] = l.data[j], l.data[i]
	return nil
}
