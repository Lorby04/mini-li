package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	lis "../lis"
)

var totalTargets = 2000000
var types = []string{"From", "To", "PAI", "Location"}

func generateTargets() {
	digit := 1000000001

	fmt.Println("Number of types:", len(types), "starting from:", digit)

	for i := 0; i < totalTargets; i++ {
		if digit%10 == 0 {
			digit++
		}

		t := lis.NewTarget(
			strconv.Itoa(digit),
			types[i%len(types)],
		)

		lis.AddTarget(t)
		if i%(totalTargets/200) == 0 {
			fmt.Println(i, "entries added")
		}
	}

	fmt.Println("End generation at:", digit)
}

func perfTest(n int) {
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

				tl := lis.NewTarget(strconv.Itoa(low), types[i%len(types)])
				rwg.Add(1)
				go func() {
					defer rwg.Done()
					lis.Query(tl)
				}()

				if high%10 == 0 {
					high--
				}

				th := lis.NewTarget(strconv.Itoa(high), types[i%len(types)])
				rwg.Add(1)
				go func() {
					defer rwg.Done()
					lis.Query(th)
				}()
			}
			rwg.Wait()
			rend := time.Now()

			fmt.Println("Round ", round, "end at:", rend, "time used:", rend.Sub(rstart))
		}(round)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println("Testing is end at: ", end, "time used:", end.Sub(start))
}