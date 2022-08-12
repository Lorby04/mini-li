package lis_tree

import (
	"fmt"
	. "mini-li/lis/target"
	"strconv"
	"sync"
	"time"
)

var totalTargets = 10000000
var types = []string{"From", "To", "PAI", "Location"}

func GenerateTargets(n int) {
	if n >= 1000 {
		totalTargets = n
	}
	digit := 1000000001

	// Gramine takes much more time on the first entry, exclude it from performance calculation
	t := NewTarget(
		strconv.Itoa(digit),
		types[0],
	)
	AddTarget(t)
	digit++

	start := time.Now()

	fmt.Println("Number of types:", len(types), "starting from:", digit, " at ", start)

	for i := 1; i < totalTargets; i++ {
		if digit%10 == 0 {
			digit++
		}
		t := NewTarget(
			strconv.Itoa(digit),
			types[i%len(types)],
		)

		AddTarget(t)
		//if i%(totalTargets/200) == 0 {
		//	fmt.Println(i, "entries added")
		//}
		digit++
	}

	fmt.Println("Writing entries:", SizeOfTargets(), ", time: ", time.Since(start))
}

func PerfTest(n, r int) {
	low := 1000000001
	high := 9999999999
	rl := r
	start := time.Now()
	fmt.Println("Start testing at:", start)

	var wg sync.WaitGroup
	for round := 0; round < n; round++ {
		f := func(round int) {
			defer wg.Done()

			var rwg sync.WaitGroup
			rstart := time.Now()
			fmt.Println("Round ", round, "start at:", rstart)
			for i := 0; i < totalTargets; i++ {
				if low%10 == 0 {
					low++
				}

				for _, ty := range types {
					t := NewTarget(strconv.Itoa(low), ty)
					rwg.Add(1)
					Query(t, func(bool) { rwg.Done() })
				}
				if high%10 == 0 {
					high--
				}

				for _, ty := range types {
					t := NewTarget(strconv.Itoa(high), ty)
					rwg.Add(1)
					Query(t, func(bool) { rwg.Done() })
				}
				low++
				high--
			}
			rwg.Wait()
			//rend := time.Now()

			//fmt.Println("Round ", round, "end at:", rend, "time used:", rend.Sub(rstart))
		}

		if r <= 1 {
			wg.Add(1)
			f(round)
		} else {
			var gwg sync.WaitGroup
			for ; round < n && round < rl; round++ {
				gwg.Add(1)
				go func() {
					defer gwg.Done()
					wg.Add(1)
					f(round)
				}()
			}
			gwg.Wait()

			if round < n {
				round--
				rl += r
			}
		}
	}
	wg.Wait()
	end := time.Now()
	fmt.Println("Searching time:", end.Sub(start), "Statistics: ", Statistics())
	Stop()
}
