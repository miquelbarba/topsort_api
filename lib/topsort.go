package lib

import (
	"github.com/dominikbraun/graph"
)

func Topsort(array [][]string) []string {
	g := graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles())

	for i := 0; i < len(array); i++ {
		_ = g.AddVertex(array[i][0])
		_ = g.AddVertex(array[i][1])

		_ = g.AddEdge(array[i][0], array[i][1])
	}

	topsort, _ := graph.TopologicalSort(g)

	return topsort
}
