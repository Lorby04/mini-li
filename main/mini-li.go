package main

import (
	"fmt"
	"os"
	"strconv"

	lm "mini-li/lis/lis_map"
	lt "mini-li/lis/lis_tree"
)

const (
	usage = "Usage: program-name testmap|testtree -n decimal(total loops:1-10000) -r decimal(number of input goroutine) -s decimal(service routines)  -t [1:1000000]M|m|K|k"
	np    = 2
	rp    = 4
	sp    = 6
	tp    = 8
)

func main() {
	args := os.Args[1:]
	if len(args) != 9 || args[np-1] != "-n" || args[tp-1] != "-t" {
		fmt.Println(usage)
		return
	}

	n, _ := strconv.Atoi(args[np])
	if n < 1 || n > 10000 {
		fmt.Println(usage)
		return
	}

	mk := func() int {
		if args[tp][len(args[tp])-1] == 'M' || args[tp][len(args[tp])-1] == 'm' {
			return 1000000
		} else if args[tp][len(args[tp])-1] == 'K' || args[4][len(args[tp])-1] == 'k' {
			return 1000
		} else {
			return 1
		}
	}()
	t, _ := strconv.Atoi(args[tp][:len(args[tp])-1])
	if t < 1 || t > 100000 {
		t = 10
	}

	r, _ := strconv.Atoi(args[rp])
	s, _ := strconv.Atoi(args[sp])

	fmt.Println("Testing started with parameters:entries:", t*mk, "query loops:", n, "input query routines:", r,
		"service routines:", s, "storage:", args[0])
	if args[0] == "testmap" {
		lm.Start(s)
		lm.GenerateTargets(t * mk)
		lm.PerfTest(n, r)
	} else {
		lt.Start(s)
		lt.GenerateTargets(t * mk)
		lt.PerfTest(n, r)
	}

	return
}
