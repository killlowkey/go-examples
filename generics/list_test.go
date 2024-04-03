package generics

import "testing"

func TestNewList(t *testing.T) {
	l := NewList[int]()
	if l == nil {
		t.Error("NewList() returned nil")
	}
}

func TestList_Append(t *testing.T) {
	l := NewList[int]()
	l.Append(1)
	if l.Len() != 1 {
		t.Errorf("Len() returned %d, want 1", l.Len())
	}
}

func TestList_Get(t *testing.T) {
	l := NewList[int]()
	l.Append(1)
	v, _ := l.Get(0)
	if v != 1 {
		t.Errorf("Get(0) returned %d, want 1", v)
	}
}

func TestList_Get_OutOfRange(t *testing.T) {
	l := NewList[int]()
	_, err := l.Get(0)
	if err == nil {
		t.Error("Get(0) did not return an error")
	}
}

func TestList_Swap(t *testing.T) {
	l := NewList[int]()
	l.Append(1)
	l.Append(2)
	l.Swap(0, 1)
	v, _ := l.Get(0)
	if v != 2 {
		t.Errorf("Get(0) returned %d, want 2", v)
	}
}

func BenchmarkList_Append(b *testing.B) {
	l := NewList[int]()
	for i := 0; i < b.N; i++ {
		l.Append(i)
	}
}

func BenchmarkList_Append_Pointer(b *testing.B) {
	type Person struct {
		Name string
	}

	l := NewList[*Person]()
	for i := 0; i < b.N; i++ {
		l.Append(&Person{Name: "John"})
	}
}
