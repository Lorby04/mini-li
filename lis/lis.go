package lis

import (
	"strings"
	"sync"
)

var lock = &sync.RWMutex{}
var targets map[Target]struct{}

func init() {
	targets = make(map[Target]struct{})
}

func wLock(lock *sync.RWMutex, locked *bool) {
	if *locked {
		return
	}
	lock.Lock()
	*locked = true
}

func wUnlock(lock *sync.RWMutex, locked *bool) {
	if !*locked {
		return
	}
	lock.Unlock()
	*locked = false
}

func rLock(lock *sync.RWMutex, locked *bool) {
	if *locked {
		return
	}
	lock.RLock()
	*locked = true
}

func rUnlock(lock *sync.RWMutex, locked *bool) {
	if !*locked {
		return
	}
	lock.RUnlock()
	*locked = false
}

func AddTarget(target Target) error {
	var key strings.Builder
	var ty strings.Builder
	key.WriteString(target.key)
	ty.WriteString(target.ty)

	t := Target{
		key: key.String(),
		ty:  ty.String(),
	}

	locked := false
	wLock(lock, &locked)
	defer func() {
		wUnlock(lock,
			&locked)
	}()

	targets[t] = struct{}{}
	wUnlock(lock, &locked)
	return nil
}

func Query(target Target) (inlist bool) {
	ch := make(chan bool)
	if ch == nil {
		panic("query")
	}

	go func() {
		locked := false
		rLock(lock, &locked)
		defer rUnlock(lock, &locked)
		_, ok := targets[target]
		rUnlock(lock, &locked)
		ch <- ok
	}()

	inlist = <-ch
	return
}
