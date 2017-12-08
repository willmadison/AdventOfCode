package advent2017

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseProgramTowerDataPoint(t *testing.T) {
	cases := []struct {
		given    string
		expected Program
	}{
		{
			"pbga (66)",
			Program{
				Name:   "pbga",
				Weight: 66,
			},
		},
		{
			"fwft (72) -> ktlj, cntj, xhth",
			Program{
				Name:   "fwft",
				Weight: 72,
				Supporting: map[string]struct{}{
					"ktlj": struct{}{},
					"cntj": struct{}{},
					"xhth": struct{}{},
				},
			},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, ParseProgramTowerDataPoint(c.given))
	}
}

func TestDeriveCallTree(t *testing.T) {
	actual := FindRootOfCallTree(strings.NewReader(`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`))

	assert.Equal(t, "tknk", actual.Name)
	assert.NotEmpty(t, actual.Subroutines)

	fmt.Println("actual:", actual)
	assert.Equal(t, 41+251+243*2, actual.TotalWeight())
}

func TestFindImbalance(t *testing.T) {
	root := FindRootOfCallTree(strings.NewReader(`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`))

	correctWeight := FindImbalance(root)

	assert.Equal(t, 60, correctWeight)
}
