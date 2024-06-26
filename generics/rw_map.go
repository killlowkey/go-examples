package generics

import (
	"errors"
	"sync"
)

type RwMap[K comparable, V any] struct {
	rwLock sync.RWMutex
	data   map[K]V
}

func NewRwMap[K comparable, V any]() *RwMap[K, V] {
	return &RwMap[K, V]{
		rwLock: sync.RWMutex{},
		data:   make(map[K]V),
	}
}

func (m *RwMap[K, V]) Insert(key K, value V) V {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	res := m.data[key]
	m.data[key] = value
	return res
}

func (m *RwMap[K, V]) Get(key K) (V, bool) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	value, ok := m.data[key]
	return value, ok
}

func (m *RwMap[K, V]) Delete(key K) V {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	res := m.data[key]
	delete(m.data, key)
	return res
}

func (m *RwMap[K, V]) Len() int {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	return len(m.data)
}

func (m *RwMap[K, V]) Range(f func(key K, value V) bool) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

func (m *RwMap[K, V]) Clear() {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	m.data = make(map[K]V)
}

func (m *RwMap[K, V]) Keys() []K {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

func (m *RwMap[K, V]) Values() []V {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	values := make([]V, 0, len(m.data))
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

func (m *RwMap[K, V]) Clone() *RwMap[K, V] {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	newMap := NewRwMap[K, V]()
	for k, v := range m.data {
		newMap.Insert(k, v)
	}
	return newMap
}

func (m *RwMap[K, V]) Merge(other *RwMap[K, V]) error {
	if other == nil {
		return errors.New("other map is nil")
	}

	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	other.rwLock.RLock()
	defer other.rwLock.RUnlock()

	other.Range(func(k K, v V) bool {
		m.Insert(k, v)
		return true
	})
	return nil
}
