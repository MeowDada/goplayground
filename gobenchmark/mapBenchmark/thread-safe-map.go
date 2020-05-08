package concurmap

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	core    map[interface{}]interface{}
	locker *sync.RWMutex
}

func New() ConcurrentMap{
	return ConcurrentMap{
		core: make(map[interface{}]interface{}),
		locker: &sync.RWMutex{},
	}
}

func (cm ConcurrentMap) Store(key, value interface{}) {
	cm.lock()
	cm.core[key] = value
	cm.unlock()
}

func (cm ConcurrentMap) Delete(key interface{}) {
	cm.lock()
	delete(cm.core, key)
	cm.unlock()
}

func (cm ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	cm.rlock()
	value, ok = cm.core[key]
	cm.runlock()
	return value, ok
}

func (cm ConcurrentMap) Len() int {
	cm.rlock()
	defer cm.runlock()
	return len(cm.core)
}

func (cm ConcurrentMap) Dump() {
	cm.rlock()
	defer cm.runlock()
	fmt.Println(cm.core)
}

func (cm ConcurrentMap) lock() {
	cm.locker.Lock()
}

func (cm ConcurrentMap) unlock() {
	cm.locker.Unlock()
}

func (cm ConcurrentMap) rlock() {
	cm.locker.RLock()
}

func (cm ConcurrentMap) runlock() {
	cm.locker.RUnlock()
}