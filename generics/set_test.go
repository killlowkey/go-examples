package generics

import "testing"

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	if s == nil {
		t.Error("NewSet() returned nil")
	}
}

func TestNewSetFromMap(t *testing.T) {
	m := map[int]any{1: nil}
	s := NewSetFromMap(m)
	if !s.Contains(1) {
		t.Error("Contains() returned false, want true")
	}
}

func TestSet_Insert(t *testing.T) {
	s := NewSet[int]()
	s.Insert(1)
	if !s.Contains(1) {
		t.Error("Contains() returned false, want true")
	}
}

func TestSet_Delete(t *testing.T) {
	s := NewSet[int]()
	s.Insert(1)
	s.Delete(1)
	if s.Contains(1) {
		t.Error("Contains() returned true, want false")
	}
}

func BenchmarkSet_Insert(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < b.N; i++ {
		s.Insert(i)
	}
}

func BenchmarkSet_Contains(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < b.N; i++ {
		s.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(i)
	}
}

func BenchmarkSet_Delete(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < b.N; i++ {
		s.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Delete(i)
	}
}
