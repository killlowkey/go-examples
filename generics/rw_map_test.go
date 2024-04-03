package generics

import "testing"

// cd generics && go test .

func TestNewRwMap(t *testing.T) {
	m := NewRwMap[int, string]()
	if m == nil {
		t.Error("NewRwMap() returned nil")
	}
}

func TestRwMap_Insert(t *testing.T) {
	m := NewRwMap[int, string]()
	res := m.Insert(1, "one")
	if res != "" {
		t.Errorf("Insert() returned %v, want %v", res, "")
	}

	res = m.Insert(1, "uno")
	if res != "one" {
		t.Errorf("Insert() returned %v, want %v", res, "one")
	}
}

func TestRwMap_Get(t *testing.T) {
	m := NewRwMap[int, string]()
	m.Insert(1, "one")
	value, ok := m.Get(1)
	if !ok {
		t.Error("Get() returned false, want true")
	}
	if value != "one" {
		t.Errorf("Get() returned %v, want %v", value, "one")
	}

	value, ok = m.Get(2)
	if ok {
		t.Error("Get() returned true, want false")
	}
	if value != "" {
		t.Errorf("Get() returned %v, want %v", value, "")
	}
}
