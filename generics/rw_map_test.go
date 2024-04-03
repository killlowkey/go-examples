package generics

import (
	"sort"
	"testing"
)

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

func TestRwMap_Delete(t *testing.T) {
	m := NewRwMap[int, string]()
	m.Insert(1, "one")
	res := m.Delete(1)
	if res != "one" {
		t.Errorf("Delete() returned %v, want %v", res, "one")
	}

	value, ok := m.Get(1)
	if ok {
		t.Error("Get() returned true, want false")
	}
	if value != "" {
		t.Errorf("Get() returned %v, want %v", value, "")
	}
}

func TestRwMap_Len(t *testing.T) {
	m := NewRwMap[int, string]()
	m.Insert(1, "one")
	m.Insert(2, "two")
	m.Insert(3, "three")
	if m.Len() != 3 {
		t.Errorf("Len() returned %v, want %v", m.Len(), 3)
	}
}

func TestRwMap_Range(t *testing.T) {
	m := NewRwMap[int, string]()
	m.Insert(1, "one")
	m.Insert(2, "two")
	m.Insert(3, "three")

	count := 0
	m.Range(func(key int, value string) bool {
		count++
		return true
	})
	if count != 3 {
		t.Errorf("Range() called %v times, want %v", count, 3)
	}

	count = 0
	m.Range(func(key int, value string) bool {
		count++
		return false
	})
	if count != 1 {
		t.Errorf("Range() called %v times, want %v", count, 1)
	}
}

func TestRwMap_Clear(t *testing.T) {
	m := NewRwMap[int, string]()
	m.Insert(1, "one")
	m.Insert(2, "two")
	m.Insert(3, "three")

	m.Clear()
	if m.Len() != 0 {
		t.Errorf("Clear() did not clear the map")
	}
}

func TestRwMap_Keys(t *testing.T) {
	m := NewRwMap[int, string]()
	m.Insert(1, "one")
	m.Insert(2, "two")
	m.Insert(3, "three")

	keys := m.Keys()
	if len(keys) != 3 {
		t.Errorf("Keys() returned %v keys, want %v", len(keys), 3)
	}

	// 对 keys 进行排序, 然后比较
	sort.Ints(keys)

	for i, key := range keys {
		if key != i+1 {
			t.Errorf("Keys() returned key %v, want %v", key, i+1)
		}
	}
}

func TestRwMap_Values(t *testing.T) {
	m := NewRwMap[int, string]()
	m.Insert(1, "one")
	m.Insert(2, "two")
	m.Insert(3, "three")

	values := m.Values()
	if len(values) != 3 {
		t.Errorf("Values() returned %v values, want %v", len(values), 3)
	}
}

func TestRwMap_Clone(t *testing.T) {
	m := NewRwMap[int, string]()
	m.Insert(1, "one")
	m.Insert(2, "two")
	m.Insert(3, "three")

	clone := m.Clone()
	if clone.Len() != 3 {
		t.Errorf("Clone() returned map with %v entries, want %v", clone.Len(), 3)
	}

	clone.Delete(1)
	if m.Len() != 3 {
		t.Errorf("Clone() did not create a deep copy")
	}
}

func TestRwMap_Concurrent(t *testing.T) {
	m := NewRwMap[int, string]()
	for i := 0; i < 1000; i++ {
		go func(i int) {
			m.Insert(i, "value")
		}(i)
	}
}
