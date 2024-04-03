package generics

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

func NewSetFromMap[T comparable](m map[T]any) *Set[T] {
	s := NewSet[T]()
	for k := range m {
		s.Insert(k)
	}
	return s
}

func (s *Set[T]) Insert(value T) {
	s.data[value] = struct{}{}
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.data[value]
	return ok
}

func (s *Set[T]) Delete(value T) {
	delete(s.data, value)
}
