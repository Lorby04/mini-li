package main

import (
	"fmt"
	"os"
	"strconv"

	lm "mini-li/lis/lis_map"
	lt "mini-li/lis/lis_tree"
)

const usage = "Usage: program-name testmap|testtree -n val[1:100]"

func main() {
	args := os.Args[1:]
	if len(args) != 3 || args[1] != "-n" {
		fmt.Println(usage)
		return
	}

	n, _ := strconv.Atoi(args[2])
	if n < 1 || n > 100 {
		fmt.Println(usage)
		return
	}
	if args[0] == "testmap" {
		lm.GenerateTargets()
		lm.PerfTest(n)
	} else {
		lt.GenerateTargets()
		lt.PerfTest(n)
	}

	return
}
