package lis_tree

import (
	"fmt"
	. "mini-li/lis/target"
	"sync"
	"sync/atomic"

	. "github.com/emirpasic/gods/maps/treemap"
)

var lock = &sync.RWMutex{}
var targets *Map = nil //map[string]struct{}
var queryAttempt uint64 = 0
var inlistAttempt uint64 = 0

var queueSize = 1000 //
var serviceNum = 100

type service struct {
	doIt     func(string)
	argument string
}

var sch chan service
var swg sync.WaitGroup

func init() {
	targets = NewWithStringComparator() //make(map[string]struct{})
	serviceNum = 1                      //runtime.NumCPU() * 4
	queueSize = serviceNum * 10
	sch = make(chan service, queueSize)
	for i := 0; i < serviceNum; i++ {
		swg.Add(1)
		go func() {
			defer swg.Done()
			var s service = service{nil, ""}
			ok := true
			for ok {
				s, ok = <-sch
				if !ok {
					return
				}

				s.doIt(s.argument)
			}
		}()
	}
}

func Stop() {
	close(sch)
	swg.Wait()
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
	//rLock(lock, &locked)
	//defer rUnlock(lock, &locked)
	//if _, found := targets.Get(target.String()); !found {
	//	rUnlock(lock, &locked)
	//	fmt.Println("Cannot find the key:\"", target.String(), "\" just inserted")
	//	panic(nil)
	//}
	return nil
}

func traverseAgainst(t string) {
	allow := false
	if !allow {
		return
	}
	fmt.Println("Query:\t\"", t, "\"")
	locked := false
	rLock(lock, &locked)
	defer rUnlock(lock, &locked)
	it := targets.Iterator()
	it.Begin()
	for more := it.Next(); more; more = it.Next() {
		fmt.Println("Valid:\t", it.Key())
	}
}

func Query(target Target, done func(bool)) {
	atomic.AddUint64(&queryAttempt, 1)
	s := service{
		doIt: func(key string) {
			locked := false
			ok := false
			defer done(ok)
			rLock(lock, &locked)
			defer rUnlock(lock, &locked)
			_, ok = targets.Get(key)
			rUnlock(lock, &locked)
			if ok {
				atomic.AddUint64(&inlistAttempt, 1)
			} else {
				traverseAgainst(key)
			}
		},
		argument: target.String(),
	}

	//sch <- s
	//go
	s.doIt(s.argument)
	return
}

func SizeOfTargets() int {
	locked := false
	rLock(lock, &locked)
	size := targets.Size() //len(targets)
	rUnlock(lock, &locked)
	return size
}
