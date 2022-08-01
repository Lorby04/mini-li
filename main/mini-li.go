package main

import (
	"fmt"
	"os"
	"strconv"
)

const usage = "Usage: program-name -n val[1:100]"

func main() {
	args := os.Args[1:]
	if len(args) != 2 || args[0] != "-n" {
		fmt.Println(usage)
		return
	}

	n, _ := strconv.Atoi(args[1])
	if n < 1 || n > 100 {
		fmt.Println(usage)
		return
	}
	generateTargets()
	perfTest(n)

	return
}
