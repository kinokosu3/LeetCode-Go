package main

import (
	"fmt"
	"sort"
)

type PointLine struct {
	Start int
	End   int
	Cost  int
}
type PointLineList []PointLine

func (p PointLineList) Len() int {
	return len(p)
}
func (p PointLineList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PointLineList) Less(i, j int) bool {
	return p[i].Cost < p[j].Cost
}
func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
func minimumEffortPath(heights [][]int) int {
	m := len(heights)
	n := len(heights[0])
	minCost := 0
	edges := make(PointLineList, 0)
	parent := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+1 < m {
				edge := PointLine{
					Start: i*n + j,
					End:   (i+1)*n + j,
					Cost:  int(abs(heights[i+1][j] - heights[i][j])),
				}
				edges = append(edges, edge)
			}
			if j+1 < n {
				edge := PointLine{
					Start: i*n + j,
					End:   i*n + j + 1,
					Cost:  int(abs(heights[i][j+1] - heights[i][j])),
				}
				edges = append(edges, edge)
			}
		}
	}
	for i := 0; i < m*n; i++ {
		parent = append(parent, i)
	}
	// 找根节点
	Find := func(x int) int {
		for {
			if parent[x] == x {
				break
			}
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	Union := func(p, q int) {
		rootP := Find(p)
		rootQ := Find(q)
		if rootP == rootQ {
			return
		}
		parent[rootP] = rootQ

	}
	isSame := func(p, q int) bool {
		return Find(p) == Find(q)
	}

	sort.Sort(edges)
	for i := 0; i < edges.Len(); i++ {
		minCost = edges[i].Cost
		Union(edges[i].Start, edges[i].End)
		if isSame(0, m*n-1) {
			break
		}
	}
	return minCost
}
func main() {
	a := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	b := minimumEffortPath(a)
	fmt.Println(b)
}
