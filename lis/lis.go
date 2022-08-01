package lis

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var lock = &sync.RWMutex{}
var targets map[string]struct{}
var queryAttempt uint64 = 0
var inlistAttempt uint64 = 0

func init() {
	targets = make(map[string]struct{})
}

func Statistics() string {
	return fmt.Sprintf("Query:%ld, Got:%ld, in %d",
		atomic.LoadUint64(&queryAttempt),
		atomic.LoadUint64(&inlistAttempt))
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
	locked := false
	wLock(lock, &locked)
	defer func() {
		wUnlock(lock,
			&locked)
	}()

	targets[target.String()] = struct{}{}
	wUnlock(lock, &locked)
	return nil
}

func Query(target Target) (inlist bool) {
	ch := make(chan bool)
	if ch == nil {
		panic("query")
	}

	atomic.AddUint64(&queryAttempt, 1)
	go func(key string) {
		locked := false
		rLock(lock, &locked)
		defer rUnlock(lock, &locked)
		_, ok := targets[key]
		rUnlock(lock, &locked)
		ch <- ok
	}(target.String())

	inlist = <-ch
	if inlist {
		atomic.AddUint64(&inlistAttempt, 1)
	}
	return
}

func SizeOfTargets() int {
	locked := false
	rLock(lock, &locked)
	size := len(targets)
	rUnlock(lock, &locked)
	return size
}
