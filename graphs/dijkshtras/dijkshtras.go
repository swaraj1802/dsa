package main

import (
	"fmt"
	"math"
)

type Edge struct {
	To     int
	From   int
	Weight int
}

type Graph map[int][]Edge

func (g *Graph) Dijikshtras(visited map[int]int,start,n int) map[int]int {
	graph := *g
	visited[start]=0
	for len(visited)< n {
		for key,value := range visited {
			for _,v := range graph[key] {
				if _,ok:=visited[v.To]; !ok {
					visited[v.To]=math.MaxInt64
				}
				if value + v.Weight < visited[v.To] {
					visited[v.To] = value+v.Weight
				}
			}
		}
	}
	return visited
}

func (g *Graph) Len() int {
	graph := *g
	return len(graph)
}

func (graph *Graph) AddEdge(f, t, w int) {
	g := *graph
	g[f] = append(g[f], Edge{
		To:     t,
		From:   f,
		Weight: w,
	})
	g[t] = append(g[t], Edge{
		To:     f,
		From:   t,
		Weight: w,
	})
}


func main() {
	graph := make(Graph)
	graph.AddEdge(1, 2, 8)
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(2, 3, 7)
	graph.AddEdge(3, 4, 9)
	graph.AddEdge(4, 5, 10)
	graph.AddEdge(5, 6, 2)
	graph.AddEdge(6, 7, 1)
	graph.AddEdge(7, 8, 7)
	graph.AddEdge(7, 0, 8)
	graph.AddEdge(1, 7, 11)
	graph.AddEdge(2, 8, 2)
	graph.AddEdge(8, 6, 6)
	graph.AddEdge(2, 5, 4)
	graph.AddEdge(3, 5, 14)
	n := graph.Len()
	visited := make(map[int]int)
	fmt.Println(graph.Dijikshtras(visited,0,n))
}
