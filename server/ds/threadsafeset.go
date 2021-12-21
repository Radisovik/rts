package ds

import "sync"

type ThreadSafeSet[K comparable] struct {
	data map[K]struct{}
	lock sync.Mutex
}

func (this *ThreadSafeSet[K]) Add(k K) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.init()
	this.data[k] = struct{}{}
}

func (this *ThreadSafeSet[K]) Remove(k K) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.data, k)
}

func (this *ThreadSafeSet[K]) Contains(k K) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, ok := this.data[k]
	return ok
}

func (this *ThreadSafeSet[K]) Len() int {
	this.lock.Lock()
	defer this.lock.Unlock()
	return len(this.data)
}

func (this *ThreadSafeSet[K]) Range(f func(K) bool) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	for k := range this.data {
		if !f(k) {
			return false
		}
	}
	return true
}

func (tsm *ThreadSafeSet[K]) init() {
	if tsm.data == nil {
		tsm.data = make(map[K]struct{})
	}
}
