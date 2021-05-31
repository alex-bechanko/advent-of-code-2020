/*
Copyright © 2021 Alex Bechanko <alex.bechanko@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseFile(t *testing.T) {
	testPath := "../../inputs/day07.example01.txt"
	expected := BagGraph{
		Nodes: map[string]bool{
			"bright white": true,
			"dark olive":   true,
			"dark orange":  true,
			"dotted black": true,
			"faded blue":   true,
			"light red":    true,
			"muted yellow": true,
			"shiny gold":   true,
			"vibrant plum": true,
		},
		Edges: []BagGraphEdge{
			{Start: "light red", End: "bright white", Weight: 1},
			{Start: "light red", End: "muted yellow", Weight: 2},
			{Start: "dark orange", End: "bright white", Weight: 3},
			{Start: "dark orange", End: "muted yellow", Weight: 4},
			{Start: "bright white", End: "shiny gold", Weight: 1},
			{Start: "muted yellow", End: "shiny gold", Weight: 2},
			{Start: "muted yellow", End: "faded blue", Weight: 9},
			{Start: "shiny gold", End: "dark olive", Weight: 1},
			{Start: "shiny gold", End: "vibrant plum", Weight: 2},
			{Start: "dark olive", End: "faded blue", Weight: 3},
			{Start: "dark olive", End: "dotted black", Weight: 4},
			{Start: "vibrant plum", End: "faded blue", Weight: 5},
			{Start: "vibrant plum", End: "dotted black", Weight: 6},
		},
	}

	data, err := ParseFile(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_Solution1(t *testing.T) {
	data := BagGraph{
		Nodes: map[string]bool{
			"bright white": true,
			"dark olive":   true,
			"dark orange":  true,
			"dotted black": true,
			"faded blue":   true,
			"light red":    true,
			"muted yellow": true,
			"shiny gold":   true,
			"vibrant plum": true,
		},
		Edges: []BagGraphEdge{
			{Start: "light red", End: "bright white", Weight: 1},
			{Start: "light red", End: "muted yellow", Weight: 2},
			{Start: "dark orange", End: "bright white", Weight: 3},
			{Start: "dark orange", End: "muted yellow", Weight: 4},
			{Start: "bright white", End: "shiny gold", Weight: 1},
			{Start: "muted yellow", End: "shiny gold", Weight: 2},
			{Start: "muted yellow", End: "faded blue", Weight: 9},
			{Start: "shiny gold", End: "dark olive", Weight: 1},
			{Start: "shiny gold", End: "vibrant plum", Weight: 2},
			{Start: "dark olive", End: "faded blue", Weight: 3},
			{Start: "dark olive", End: "dotted black", Weight: 4},
			{Start: "vibrant plum", End: "faded blue", Weight: 5},
			{Start: "vibrant plum", End: "dotted black", Weight: 6},
		},
	}
	expected := "4"
	actual, err := Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

}

func Test_Solution2(t *testing.T) {
	data := BagGraph{
		Nodes: map[string]bool{
			"bright white": true,
			"dark olive":   true,
			"dark orange":  true,
			"dotted black": true,
			"faded blue":   true,
			"light red":    true,
			"muted yellow": true,
			"shiny gold":   true,
			"vibrant plum": true,
		},
		Edges: []BagGraphEdge{
			{Start: "light red", End: "bright white", Weight: 1},
			{Start: "light red", End: "muted yellow", Weight: 2},
			{Start: "dark orange", End: "bright white", Weight: 3},
			{Start: "dark orange", End: "muted yellow", Weight: 4},
			{Start: "bright white", End: "shiny gold", Weight: 1},
			{Start: "muted yellow", End: "shiny gold", Weight: 2},
			{Start: "muted yellow", End: "faded blue", Weight: 9},
			{Start: "shiny gold", End: "dark olive", Weight: 1},
			{Start: "shiny gold", End: "vibrant plum", Weight: 2},
			{Start: "dark olive", End: "faded blue", Weight: 3},
			{Start: "dark olive", End: "dotted black", Weight: 4},
			{Start: "vibrant plum", End: "faded blue", Weight: 5},
			{Start: "vibrant plum", End: "dotted black", Weight: 6},
		},
	}
	expected := "32"
	actual, err := Solution2(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
