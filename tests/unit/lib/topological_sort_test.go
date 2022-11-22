package lib

import (
	"reflect"
	"testing"

	"github.com/miquelbarba/topsort_api/lib"
)

func TestTopologicalSort(t *testing.T) {
	input := [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}
	expected := []string{"SFO", "ATL", "GSO", "IND", "EWR"}
	testTopologicalSort(t, input, expected)

	input = [][]string{{"SFO", "EWR"}}
	expected = []string{"SFO", "EWR"}
	testTopologicalSort(t, input, expected)

	input = [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}
	expected = []string{"SFO", "ATL", "EWR"}
	testTopologicalSort(t, input, expected)
}

func TestTopologicalSortEmpty(t *testing.T) {
	input := [][]string{}
	expected := []string{}
	testTopologicalSort(t, input, expected)
}

func testTopologicalSort(t *testing.T, input [][]string, expected []string) {
	t.Helper()

	result := lib.TopologicalSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %q, expected %q", result, expected)
	}
}
