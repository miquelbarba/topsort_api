package lib

import (
	"errors"
	"fmt"

	"github.com/dominikbraun/graph"
)

func TopologicalSort(edges [][]string) ([]string, error) {
	if len(edges) == 0 {
		return []string{}, nil
	}

	uniqueEdges := uniqueEdges(edges)
	topsort, err := graph.TopologicalSort(buildGraph(uniqueEdges))

	if err != nil {
		return nil, err
	}

	if len(uniqueEdges)+1 != len(topsort) {
		return nil, errors.New("graph with separated paths")
	}

	return topsort, nil
}

func buildGraph(edges [][]string) graph.Graph[string, string] {
	g := graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles())

	for _, vertex := range uniqueVertices(edges) {
		_ = g.AddVertex(vertex)
	}

	for _, edge := range edges {
		_ = g.AddEdge(edge[0], edge[1])
	}

	return g
}

func uniqueVertices(edges [][]string) []string {
	return unique(flatten(edges))
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

func uniqueEdges(s [][]string) [][]string {
	inResult := make(map[string]bool)

	var result [][]string

	for _, item := range s {
		str := fmt.Sprintf("%s_%s", item[0], item[1])
		if _, ok := inResult[str]; !ok {
			inResult[str] = true

			result = append(result, item)
		}
	}

	return result
}
