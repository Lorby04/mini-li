package main

import (
	"fmt"
	"os"
	"strconv"

	lm "mini-li/lis/lis_map"
	lt "mini-li/lis/lis_tree"
)

const usage = "Usage: program-name testmap|testtree -n val[1:100] -t [1:1000]M|m|K|k"

func main() {
	args := os.Args[1:]
	if len(args) != 5 || args[1] != "-n" || args[3] != "-t" {
		fmt.Println(usage)
		return
	}

	n, _ := strconv.Atoi(args[2])
	if n < 1 || n > 100 {
		fmt.Println(usage)
		return
	}

	mk := func() int {
		if args[4][len(args[4])-1] == 'M' || args[4][len(args[4])-1] == 'm' {
			return 1000000
		} else if args[4][len(args[4])-1] == 'K' || args[4][len(args[4])-1] == 'k' {
			return 1000
		} else {
			return 1
		}
	}()
	t, _ := strconv.Atoi(args[4][:len(args[4])-1])
	if t < 1 || t > 1000 {
		t = 10
	}

	if args[0] == "testmap" {
		lm.GenerateTargets(t * mk)
		lm.PerfTest(n)
	} else {
		lt.GenerateTargets(t * mk)
		lt.PerfTest(n)
	}

	return
}
