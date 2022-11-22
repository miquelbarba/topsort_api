package lib

import (
	"fmt"
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

func TestTopologicalSortError(t *testing.T) {
	input := [][]string{{"IND", "EWR"}, {"SFO", "ATL"}}
	_, err := lib.TopologicalSort(input)

	if err.Error() != "invalid graph" {
		t.Errorf("expected to receive an error")
	}
}

func TestTopologicalSortErrorWithRepeated(t *testing.T) {
	input := [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"SFO", "ATL"}}
	res, err := lib.TopologicalSort(input)

	fmt.Println(res)
	fmt.Println(err)

	if err.Error() != "invalid graph" {
		t.Errorf("expected to receive an error")
	}
}

func testTopologicalSort(t *testing.T, input [][]string, expected []string) {
	t.Helper()

	result, _ := lib.TopologicalSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %q, expected %q", result, expected)
	}
}
