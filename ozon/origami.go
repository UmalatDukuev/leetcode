package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ x, y int }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, m, k int
		fmt.Fscan(in, &n, &m, &k)
		grid := make([][]rune, n)
		for i := 0; i < n; i++ {
			row := make([]rune, m)
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &row[j])
			}
			grid[i] = row
		}
		border := func(idx int) P {
			per := 2*(n+m) - 4
			idx = (idx - 1) % per
			if idx < m {
				return P{0, idx}
			}
			idx -= m
			if idx < n-1 {
				return P{idx + 1, m - 1}
			}
			idx -= n - 1
			if idx < m-1 {
				return P{n - 1, m - 2 - idx}
			}
			idx -= m - 1
			return P{n - 2 - idx, 0}
		}
		for op := 0; op < k; op++ {
			var a, b int
			fmt.Fscan(in, &a, &b)
			A := border(a)
			B := border(b)
			dx := B.x - A.x
			dy := B.y - A.y
			norm := dx*dx + dy*dy
			newGrid := make([][]rune, n)
			for i := range newGrid {
				newGrid[i] = make([]rune, m)
				for j := range newGrid[i] {
					newGrid[i][j] = '.'
				}
			}
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					if grid[i][j] == '.' {
						continue
					}
					vx := i - A.x
					vy := j - A.y
					cross := dx*vy - dy*vx
					side := cross > 0
					nx, ny := i, j
					if side {
						dot := vx*dx + vy*dy
						projx := A.x + (dot*dx)/norm
						projy := A.y + (dot*dy)/norm
						rx := projx*2 - i
						ry := projy*2 - j
						nx, ny = rx, ry
					}
					if 0 <= nx && nx < n && 0 <= ny && ny < m {
						if grid[i][j] == '#' {
							newGrid[nx][ny] = '#'
						} else {
							newGrid[nx][ny] = grid[i][j]
						}
					}
				}
			}
			grid = newGrid
			minx, maxx, miny, maxy := n, 0, m, 0
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					if grid[i][j] != '.' {
						if i < minx {
							minx = i
						}
						if i > maxx {
							maxx = i
						}
						if j < miny {
							miny = j
						}
						if j > maxy {
							maxy = j
						}
					}
				}
			}
			for i := minx; i <= maxx; i++ {
				for j := miny; j <= maxy; j++ {
					out.WriteRune(grid[i][j])
				}
				out.WriteByte('\n')
			}
			out.WriteByte('\n')
		}
	}
}
