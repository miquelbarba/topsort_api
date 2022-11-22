package main

import (
	"fmt"

	"github.com/miquelbarba/topsort_api/lib"
)

func main() {
	array := [][]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}

	sort := lib.Topsort(array)
	fmt.Println(sort)
}
