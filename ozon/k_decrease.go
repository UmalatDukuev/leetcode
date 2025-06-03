package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(in, &n, &k)
		a := make([]int64, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		expire := make([]int64, n+1)
		var applied int64
		ok := true
		for i := 0; i < n; i++ {
			applied -= expire[i]
			need := a[i] - applied
			if need > 0 {
				if i+k > n {
					ok = false
					break
				}
				applied += need
				expire[i+k] += need
			}
			if need < 0 {
				ok = false
				break
			}
		}
		if ok {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
