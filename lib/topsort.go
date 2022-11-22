package lib

import (
	"github.com/dominikbraun/graph"
)

func Topsort(edges [][]string) []string {
	topsort, _ := graph.TopologicalSort(buildGraph(edges))

	return topsort
}

func buildGraph(edges [][]string) graph.Graph[string, string] {
	g := graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles())

	vertices := unique(flatten(edges))
	for _, vertex := range vertices {
		_ = g.AddVertex(vertex)
	}

	for _, edge := range edges {
		_ = g.AddEdge(edge[0], edge[1])
	}

	return g
}

func flatten[T any](lists [][]T) []T {
	var result []T
	for _, list := range lists {
		result = append(result, list...)
	}

	return result
}

func unique(s []string) []string {
	inResult := make(map[string]bool)

	var result []string

	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true

			result = append(result, str)
		}
	}

	return result
}
