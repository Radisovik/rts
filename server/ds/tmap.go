package ds

import "sync"

type ThreadSafeMap[K comparable, V any] struct {
	data map[K]V
	lock sync.Mutex
}

func (tsm *ThreadSafeMap[K, V]) Put(k K, v V) {
	tsm.lock.Lock()
	defer tsm.lock.Unlock()
	tsm.init()
	tsm.data[k] = v
}
func (tsm *ThreadSafeMap[K, V]) Get(k K) (V, bool) {
	tsm.lock.Lock()
	defer tsm.lock.Unlock()
	if tsm.data == nil {
		var z V
		return z, false
	}
	return tsm.data[k], true
}
func (tsm *ThreadSafeMap[K, V]) Len() int {
	tsm.lock.Lock()
	defer tsm.lock.Unlock()
	return len(tsm.data)
}

func (tsm *ThreadSafeMap[K, V]) Clear() {
	tsm.lock.Lock()
	defer tsm.lock.Unlock()
	tsm.data = nil
}

func (tsm *ThreadSafeMap[K, V]) init() {
	if tsm.data == nil {
		tsm.data = make(map[K]V)
	}
}
