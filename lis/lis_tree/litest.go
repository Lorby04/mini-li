package lis_tree

import (
	"fmt"
	. "mini-li/lis/target"
	"strconv"
	"sync"
	"time"
)

var totalTargets = 2000000
var types = []string{"From", "To", "PAI", "Location"}

func GenerateTargets() {
	digit := 1000000001

	// Gramine takes much more time on the first entry, exclude it from performance calculation
	t := NewTarget(
		strconv.Itoa(digit),
		types[0],
	)
	AddTarget(t)
	digit++

	start := time.Now()

	fmt.Println("Number of types:", len(types), "starting from:", digit)

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

	fmt.Println("End generation at:", digit, "total entries:", SizeOfTargets(), "in ", time.Now().Sub(start))
}

func PerfTest(n int) {
	low := 1000000001
	high := 9999999999

	start := time.Now()
	fmt.Println("Start testing at:", start)

	var wg sync.WaitGroup
	for round := 0; round < n; round++ {
		wg.Add(1)
		go func(round int) {
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
			rend := time.Now()

			fmt.Println("Round ", round, "end at:", rend, "time used:", rend.Sub(rstart))
		}(round)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println("Testing is end at: ", end, "time used:", end.Sub(start), "Statistics: ", Statistics())
	Stop()
}
