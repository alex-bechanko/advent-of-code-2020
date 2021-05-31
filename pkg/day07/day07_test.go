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
	expected := bagGraph{
		nodes: map[string]bool{
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
		edges: []bagGraphEdge{
			{start: "light red", end: "bright white", weight: 1},
			{start: "light red", end: "muted yellow", weight: 2},
			{start: "dark orange", end: "bright white", weight: 3},
			{start: "dark orange", end: "muted yellow", weight: 4},
			{start: "bright white", end: "shiny gold", weight: 1},
			{start: "muted yellow", end: "shiny gold", weight: 2},
			{start: "muted yellow", end: "faded blue", weight: 9},
			{start: "shiny gold", end: "dark olive", weight: 1},
			{start: "shiny gold", end: "vibrant plum", weight: 2},
			{start: "dark olive", end: "faded blue", weight: 3},
			{start: "dark olive", end: "dotted black", weight: 4},
			{start: "vibrant plum", end: "faded blue", weight: 5},
			{start: "vibrant plum", end: "dotted black", weight: 6},
		},
	}

	data, err := ParseFile(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_Solution1(t *testing.T) {
	data := bagGraph{
		nodes: map[string]bool{
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
		edges: []bagGraphEdge{
			{start: "light red", end: "bright white", weight: 1},
			{start: "light red", end: "muted yellow", weight: 2},
			{start: "dark orange", end: "bright white", weight: 3},
			{start: "dark orange", end: "muted yellow", weight: 4},
			{start: "bright white", end: "shiny gold", weight: 1},
			{start: "muted yellow", end: "shiny gold", weight: 2},
			{start: "muted yellow", end: "faded blue", weight: 9},
			{start: "shiny gold", end: "dark olive", weight: 1},
			{start: "shiny gold", end: "vibrant plum", weight: 2},
			{start: "dark olive", end: "faded blue", weight: 3},
			{start: "dark olive", end: "dotted black", weight: 4},
			{start: "vibrant plum", end: "faded blue", weight: 5},
			{start: "vibrant plum", end: "dotted black", weight: 6},
		},
	}
	expected := "4"
	actual, err := Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

}

func Test_Solution2(t *testing.T) {
	data := bagGraph{
		nodes: map[string]bool{
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
		edges: []bagGraphEdge{
			{start: "light red", end: "bright white", weight: 1},
			{start: "light red", end: "muted yellow", weight: 2},
			{start: "dark orange", end: "bright white", weight: 3},
			{start: "dark orange", end: "muted yellow", weight: 4},
			{start: "bright white", end: "shiny gold", weight: 1},
			{start: "muted yellow", end: "shiny gold", weight: 2},
			{start: "muted yellow", end: "faded blue", weight: 9},
			{start: "shiny gold", end: "dark olive", weight: 1},
			{start: "shiny gold", end: "vibrant plum", weight: 2},
			{start: "dark olive", end: "faded blue", weight: 3},
			{start: "dark olive", end: "dotted black", weight: 4},
			{start: "vibrant plum", end: "faded blue", weight: 5},
			{start: "vibrant plum", end: "dotted black", weight: 6},
		},
	}
	expected := "32"
	actual, err := Solution2(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
