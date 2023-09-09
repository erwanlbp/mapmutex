package mapmutex

import (
	"fmt"
	"sync"
)

type Mutexs struct {
	mtxs    sync.Map
	mainMtx sync.Mutex
}

func (mm *Mutexs) Lock(resource string) {
	v, ok := mm.mtxs.Load(resource)
	if !ok {
		mm.mainMtx.Lock()
		defer mm.mainMtx.Unlock()

		// Double check if mutex is already present for resource in case another Lock was acquired while checking 'ok'
		if m, ok := mm.mtxs.Load(resource); ok {
			v = m
		} else {
			v = new(sync.Mutex)
			mm.mtxs.Store(resource, v)
		}
	}
	mtx, _ := v.(*sync.Mutex)
	mtx.Lock()
}

func (mm *Mutexs) Unlock(resource string) {
	v, ok := mm.mtxs.Load(resource)
	if !ok {
		panic(fmt.Errorf("resource '%s' isn't locked", resource))
	}
	mtx, _ := v.(*sync.Mutex)
	mtx.Unlock()
}
