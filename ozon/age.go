package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var parent, diff, rnk []int
var idxMap map[string]int

func getIdx(s string) int {
	if v, ok := idxMap[s]; ok {
		return v
	}
	v := len(parent)
	idxMap[s] = v
	parent = append(parent, v)
	diff = append(diff, 0)
	rnk = append(rnk, 0)
	return v
}
func find(x int) (int, int) {
	if parent[x] == x {
		return x, 0
	}
	p := parent[x]
	r, d := find(p)
	diff[x] += d
	parent[x] = r
	return r, diff[x]
}
func unite(x, y, D int) {
	rx, dx := find(x)
	ry, dy := find(y)
	if rx == ry {
		return
	}
	if rnk[rx] < rnk[ry] {
		parent[rx] = ry
		diff[rx] = dy + D - dx
	} else {
		parent[ry] = rx
		diff[ry] = dx - dy - D
		if rnk[rx] == rnk[ry] {
			rnk[rx]++
		}
	}
}
func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(line))
	for ; t > 0; t-- {
		line, _ = in.ReadString('\n')
		target := strings.TrimSuffix(strings.TrimSpace(line[len("How old is "):]), "?")
		parent = nil
		diff = nil
		rnk = nil
		idxMap = make(map[string]int)
		base := getIdx("BASE")
		ti := getIdx(target)
		for i := 0; i < 3; i++ {
			s, _ := in.ReadString('\n')
			w := strings.Fields(strings.TrimSpace(s))
			Ai := getIdx(w[0])
			if w[2] == "the" {
				unite(Ai, getIdx(w[6]), 0)
			} else {
				num, _ := strconv.Atoi(w[2])
				if w[3] == "years" {
					if w[4] == "old" {
						unite(Ai, base, num)
					} else if w[4] == "younger" {
						unite(Ai, getIdx(w[6]), -num)
					} else {
						unite(Ai, getIdx(w[6]), num)
					}
				}
			}
		}
		_, da := find(ti)
		_, db := find(base)
		fmt.Println(da - db)
	}
}
