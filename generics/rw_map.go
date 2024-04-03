package generics

import "sync"

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
