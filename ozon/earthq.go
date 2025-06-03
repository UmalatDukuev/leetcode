package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, m, x, y, p int
		fmt.Fscan(in, &n, &m)
		grid := make([][]byte, n)
		for i := 0; i < n; i++ {
			var s string
			fmt.Fscan(in, &s)
			grid[i] = []byte(s)
		}
		fmt.Fscan(in, &x, &y, &p)
		x--
		y--
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		// 4-directional connectivity for houses
		dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		ans := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] == '0' || vis[i][j] {
					continue
				}
				// BFS for this house
				q := [][2]int{{i, j}}
				vis[i][j] = true
				destroyed := false
				for k := 0; k < len(q); k++ {
					di, dj := q[k][0], q[k][1]
					// compute wave strength here
					d := p - max(abs(di-x), abs(dj-y))
					if d > 0 && int(grid[di][dj]-'0') < d {
						destroyed = true
					}
					// expand to 4-neighbors only
					for _, dr := range dirs {
						ni, nj := di+dr[0], dj+dr[1]
						if ni >= 0 && ni < n && nj >= 0 && nj < m && !vis[ni][nj] && grid[ni][nj] != '0' {
							vis[ni][nj] = true
							q = append(q, [2]int{ni, nj})
						}
					}
				}
				if destroyed {
					ans++
				}
			}
		}
		fmt.Fprintln(out, ans)
	}
}
