package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph/simple"
)

func main() {
	g := simple.NewDirectedGraph()

	n1 := g.NewNode()
	g.AddNode(n1)

	n2 := g.NewNode()
	g.AddNode(n2)

	g.SetEdge(g.NewEdge(n1, n2))

	fmt.Println("Nodes:", g.Nodes().Len())
	fmt.Println("Edges from n1:", g.From(n1.ID()).Len())
	fmt.Println(g)
}
