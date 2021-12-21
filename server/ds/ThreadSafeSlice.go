package ds

import (
	"fmt"
	"sync"
)

type ThreadSafeSlice[V comparable] struct {
	data []V
	lock sync.Mutex
}

func (this *ThreadSafeSlice[V]) String() string {
	this.lock.Lock()
	defer this.lock.Unlock()
	return fmt.Sprintf("%v", this.data)
}

func (this *ThreadSafeSlice[V]) Append(v V) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.data = append(this.data, v)
}

func (this *ThreadSafeSlice[V]) Remove(v V) {
	this.lock.Lock()
	defer this.lock.Unlock()
	var newData []V
	for _, vv := range this.data {
		if v != vv {
			newData = append(newData, vv)
		}
	}
	this.data = newData
}
func (this *ThreadSafeSlice[V]) Range(f func(V) bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for _, v := range this.data {
		if !f(v) {
			return
		}
	}
}
