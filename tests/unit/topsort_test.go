package lib

import (
	"reflect"
	"testing"

	"github.com/miquelbarba/topsort_api/lib"
)

func TestTopsort(t *testing.T) {
	input := [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}
	expected := []string{"SFO", "ATL", "GSO", "IND", "EWR"}
	_testTopSort(t, input, expected)

	input = [][]string{{"SFO", "EWR"}}
	expected = []string{"SFO", "EWR"}
	_testTopSort(t, input, expected)

	input = [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}
	expected = []string{"SFO", "ATL", "EWR"}
	_testTopSort(t, input, expected)
}

func _testTopSort(t *testing.T, input [][]string, expected []string) {
	result := lib.Topsort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %q, expected %q", result, expected)
	}
}
