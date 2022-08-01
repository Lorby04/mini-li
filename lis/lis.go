package lis

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/emirpasic/gods/maps/treemap"
)

var lock = &sync.RWMutex{}
var targets *Map = nil //map[string]struct{}
var queryAttempt uint64 = 0
var inlistAttempt uint64 = 0

func init() {
	targets = treemap.NewWithStringComparator() //make(map[string]struct{})
}

func Statistics() string {
	return fmt.Sprintf("Query:%d, Got:%d, in %d entries",
		atomic.LoadUint64(&queryAttempt),
		atomic.LoadUint64(&inlistAttempt),
		SizeOfTargets())
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

	//fmt.Println("Add:", target.String(), "before:", len(targets))
	targets.Put(target.String(), struct{}{})
	//fmt.Println("Added:", target.String(), "after:", len(targets))
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
		_, ok := targets.Get(key)
		rUnlock(lock, &locked)
		if ok {
			atomic.AddUint64(&inlistAttempt, 1)
		}
		ch <- ok
	}(target.String())

	inlist = <-ch
	return
}

func SizeOfTargets() int {
	locked := false
	rLock(lock, &locked)
	size := targets.Size() //len(targets)
	rUnlock(lock, &locked)
	return size
}
