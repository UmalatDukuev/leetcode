package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Start   = "M"
	Restart = "R"
	Cancel  = "C"
	Delete  = "D"
)

func processLegit(proc string) string {
	yes := "YES"
	no := "NO"
	proc_len := len(proc) - 1
	if string(proc[0]) != Start || string(proc[proc_len]) != Delete {
		return no
	}
	status := Start
	for i, val := range proc {
		if i == 0 {
			continue
		}
		let := string(val)
		if status == Start && let == Start {
			return no
		}

		if status == Restart && (let == Start || let == Delete || let == Restart) {
			return no
		}
		if status == Cancel && (let == Delete || let == Restart || let == Cancel) {
			return no
		}
		if status == Delete && (let == Delete || let == Cancel || let == Restart) {
			return no
		}
		status = let
	}

	return yes
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a int
	fmt.Fscan(in, &a)
	procs := make([]string, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &procs[i])
	}
	for _, val := range procs {
		fmt.Println(processLegit(val))
	}
}
