package lib

import (
	"reflect"
	"testing"

	"github.com/miquelbarba/topsort_api/lib"
)

func TestTopsort(t *testing.T) {
	input := [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}
	expected := []string{"SFO", "ATL", "GSO", "IND", "EWR"}
	testTopSort(t, input, expected)

	input = [][]string{{"SFO", "EWR"}}
	expected = []string{"SFO", "EWR"}
	testTopSort(t, input, expected)

	input = [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}
	expected = []string{"SFO", "ATL", "EWR"}
	testTopSort(t, input, expected)
}

func TestTopsortEmpty(t *testing.T) {
	input := [][]string{}
	expected := []string{}
	testTopSort(t, input, expected)
}

func testTopSort(t *testing.T, input [][]string, expected []string) {
	t.Helper()

	result := lib.Topsort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %q, expected %q", result, expected)
	}
}
