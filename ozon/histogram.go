package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for tc := 0; tc < t; tc++ {
		var N, M int
		fmt.Fscan(in, &N, &M)
		freq := make(map[string]int, N)
		var ans int64
		row := make([]int, M)
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				fmt.Fscan(in, &row[j])
			}
			base := row[0]
			last := row[M-1]
			sig := make([]string, M)
			rev := make([]string, M)
			for k := 0; k < M; k++ {
				sig[k] = fmt.Sprint(row[k] - base)
				rev[k] = fmt.Sprint(last - row[M-1-k])
			}
			revKey := strings.Join(rev, ",")
			ans += int64(freq[revKey])
			key := strings.Join(sig, ",")
			freq[key]++
		}
		fmt.Fprintln(out, ans)
	}
}
