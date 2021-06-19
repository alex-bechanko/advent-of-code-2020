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
package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	var data string
	var expected PuzzleInput
	var actual PuzzleInput
	var err error

	data = "../../inputs/day17.example.txt"
	expected = PuzzleInput{
		Cubes: map[Vertex]bool{
			{x: 1, y: 0, z: 0}: true,
			{x: 2, y: 1, z: 0}: true,
			{x: 0, y: 2, z: 0}: true,
			{x: 1, y: 2, z: 0}: true,
			{x: 2, y: 2, z: 0}: true,
		},
	}
	actual, err = Parse(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Vertex_Neighbors(t *testing.T) {
	var data Vertex
	var expected []Vertex
	var actual []Vertex

	data = Vertex{x: 0, y: 0, z: 0}
	expected = []Vertex{
		{x: -1, y: -1, z: -1},
		{x: -1, y: -1, z: 0},
		{x: -1, y: -1, z: 1},
		{x: -1, y: 0, z: -1},
		{x: -1, y: 0, z: 0},
		{x: -1, y: 0, z: 1},
		{x: -1, y: 1, z: -1},
		{x: -1, y: 1, z: 0},
		{x: -1, y: 1, z: 1},
		{x: 0, y: -1, z: -1},
		{x: 0, y: -1, z: 0},
		{x: 0, y: -1, z: 1},
		{x: 0, y: 0, z: -1},
		{x: 0, y: 0, z: 1},
		{x: 0, y: 1, z: -1},
		{x: 0, y: 1, z: 0},
		{x: 0, y: 1, z: 1},
		{x: 1, y: -1, z: -1},
		{x: 1, y: -1, z: 0},
		{x: 1, y: -1, z: 1},
		{x: 1, y: 0, z: -1},
		{x: 1, y: 0, z: 0},
		{x: 1, y: 0, z: 1},
		{x: 1, y: 1, z: -1},
		{x: 1, y: 1, z: 0},
		{x: 1, y: 1, z: 1},
	}
	actual = data.Neighbors()

	assert.ElementsMatch(t, expected, actual)
}

func Test_Vertex_NumActiveNeighbors(t *testing.T) {
	data1 := Vertex{x: 0, y: 0, z: 0}
	data2 := map[Vertex]bool{
		{x: 1, y: 0, z: 0}:   true,
		{x: 2, y: 0, z: 0}:   true,  //shouldn't count as a neighbor
		{x: 0, y: 1, z: 0}:   false, // shouldn't count because inactive
		{x: -1, y: 0, z: -1}: true,
	}

	expected := 2
	actual := data1.NumActiveNeighbors(data2)
	assert.Equal(t, expected, actual)
}

func Test_PuzzleInput_AllAdjacentVertices(t *testing.T) {
	var data PuzzleInput
	var expected map[Vertex]bool
	var actual map[Vertex]bool
	var err error

	data, err = Parse("../../inputs/day17.example.txt")
	assert.NoError(t, err)

	// basically all the neighbors of active vertices
	// I'm glad golang lets you overwrite map values within the
	// same definition, otherwise this would be less straightforward
	// to calculate
	expected = map[Vertex]bool{
		// neighbors of {x: 1, y: 0, z: 0}
		{x: 1, y: 0, z: 0}:   true,
		{x: 0, y: 0, z: 0}:   false,
		{x: 2, y: 0, z: 0}:   false,
		{x: 0, y: -1, z: 0}:  false,
		{x: 1, y: -1, z: 0}:  false,
		{x: 2, y: -1, z: 0}:  false,
		{x: 0, y: 1, z: 0}:   false,
		{x: 1, y: 1, z: 0}:   false,
		{x: 2, y: 1, z: 0}:   true,
		{x: 1, y: 0, z: -1}:  false,
		{x: 0, y: 0, z: -1}:  false,
		{x: 2, y: 0, z: -1}:  false,
		{x: 0, y: -1, z: -1}: false,
		{x: 1, y: -1, z: -1}: false,
		{x: 2, y: -1, z: -1}: false,
		{x: 0, y: 1, z: -1}:  false,
		{x: 1, y: 1, z: -1}:  false,
		{x: 2, y: 1, z: -1}:  false,
		{x: 1, y: 0, z: 1}:   false,
		{x: 0, y: 0, z: 1}:   false,
		{x: 2, y: 0, z: 1}:   false,
		{x: 0, y: -1, z: 1}:  false,
		{x: 1, y: -1, z: 1}:  false,
		{x: 2, y: -1, z: 1}:  false,
		{x: 0, y: 1, z: 1}:   false,
		{x: 1, y: 1, z: 1}:   false,
		{x: 2, y: 1, z: 1}:   false,

		// neighbors of {x: 2, y: 1, z: 0}
		{x: 1, y: 1, z: 0}:  false,
		{x: 2, y: 1, z: 0}:  true,
		{x: 3, y: 1, z: 0}:  false,
		{x: 1, y: 0, z: 0}:  true,
		{x: 2, y: 0, z: 0}:  false,
		{x: 3, y: 0, z: 0}:  false,
		{x: 1, y: 1, z: 0}:  true,
		{x: 2, y: 1, z: 0}:  true,
		{x: 3, y: 1, z: 0}:  false,
		{x: 1, y: 1, z: -1}: false,
		{x: 2, y: 1, z: -1}: false,
		{x: 3, y: 1, z: -1}: false,
		{x: 1, y: 0, z: -1}: false,
		{x: 2, y: 0, z: -1}: false,
		{x: 3, y: 0, z: -1}: false,
		{x: 1, y: 1, z: -1}: false,
		{x: 2, y: 1, z: -1}: false,
		{x: 3, y: 1, z: -1}: false,
		{x: 1, y: 1, z: 1}:  false,
		{x: 2, y: 1, z: 1}:  false,
		{x: 3, y: 1, z: 1}:  false,
		{x: 1, y: 0, z: 1}:  false,
		{x: 2, y: 0, z: 1}:  false,
		{x: 3, y: 0, z: 1}:  false,
		{x: 1, y: 1, z: 1}:  false,
		{x: 2, y: 1, z: 1}:  false,
		{x: 3, y: 1, z: 1}:  false,

		// neighbors of {x: 0, y: 2, z: 0}
		{x: -1, y: 2, z: 0}:  false,
		{x: 0, y: 2, z: 0}:   true,
		{x: 1, y: 2, z: 0}:   true,
		{x: -1, y: 1, z: 0}:  false,
		{x: 0, y: 1, z: 0}:   false,
		{x: 1, y: 1, z: 0}:   false,
		{x: -1, y: 3, z: 0}:  false,
		{x: 0, y: 3, z: 0}:   false,
		{x: 1, y: 3, z: 0}:   false,
		{x: -1, y: 2, z: -1}: false,
		{x: 0, y: 2, z: -1}:  false,
		{x: 1, y: 2, z: -1}:  false,
		{x: -1, y: 1, z: -1}: false,
		{x: 0, y: 1, z: -1}:  false,
		{x: 1, y: 1, z: -1}:  false,
		{x: -1, y: 3, z: -1}: false,
		{x: 0, y: 3, z: -1}:  false,
		{x: 1, y: 3, z: -1}:  false,
		{x: -1, y: 2, z: 1}:  false,
		{x: 0, y: 2, z: 1}:   false,
		{x: 1, y: 2, z: 1}:   false,
		{x: -1, y: 1, z: 1}:  false,
		{x: 0, y: 1, z: 1}:   false,
		{x: 1, y: 1, z: 1}:   false,
		{x: -1, y: 3, z: 1}:  false,
		{x: 0, y: 3, z: 1}:   false,
		{x: 1, y: 3, z: 1}:   false,

		//neighbors of  {x: 1, y: 2, z: 0}
		{x: 0, y: 2, z: 0}:  true,
		{x: 1, y: 2, z: 0}:  true,
		{x: 2, y: 2, z: 0}:  true,
		{x: 0, y: 1, z: 0}:  false,
		{x: 1, y: 1, z: 0}:  false,
		{x: 2, y: 1, z: 0}:  true,
		{x: 0, y: 3, z: 0}:  false,
		{x: 1, y: 3, z: 0}:  false,
		{x: 2, y: 3, z: 0}:  false,
		{x: 0, y: 2, z: -1}: false,
		{x: 1, y: 2, z: -1}: false,
		{x: 2, y: 2, z: -1}: false,
		{x: 0, y: 1, z: -1}: false,
		{x: 1, y: 1, z: -1}: false,
		{x: 2, y: 1, z: -1}: false,
		{x: 0, y: 3, z: -1}: false,
		{x: 1, y: 3, z: -1}: false,
		{x: 2, y: 3, z: -1}: false,
		{x: 0, y: 2, z: 1}:  false,
		{x: 1, y: 2, z: 1}:  false,
		{x: 2, y: 2, z: 1}:  false,
		{x: 0, y: 1, z: 1}:  false,
		{x: 1, y: 1, z: 1}:  false,
		{x: 2, y: 1, z: 1}:  false,
		{x: 0, y: 3, z: 1}:  false,
		{x: 1, y: 3, z: 1}:  false,
		{x: 2, y: 3, z: 1}:  false,
		{x: 1, y: 2, z: 0}:  true,
		{x: 2, y: 2, z: 0}:  true,
		{x: 3, y: 2, z: 0}:  false,
		{x: 1, y: 1, z: 0}:  false,
		{x: 2, y: 1, z: 0}:  true,
		{x: 3, y: 1, z: 0}:  false,
		{x: 1, y: 3, z: 0}:  false,
		{x: 2, y: 3, z: 0}:  false,
		{x: 3, y: 3, z: 0}:  false,
		{x: 1, y: 2, z: -1}: false,
		{x: 2, y: 2, z: -1}: false,
		{x: 3, y: 2, z: -1}: false,
		{x: 1, y: 1, z: -1}: false,
		{x: 2, y: 1, z: -1}: false,
		{x: 3, y: 1, z: -1}: false,
		{x: 1, y: 3, z: -1}: false,
		{x: 2, y: 3, z: -1}: false,
		{x: 3, y: 3, z: -1}: false,
		{x: 1, y: 2, z: 1}:  false,
		{x: 2, y: 2, z: 1}:  false,
		{x: 3, y: 2, z: 1}:  false,
		{x: 1, y: 1, z: 1}:  false,
		{x: 2, y: 1, z: 1}:  false,
		{x: 3, y: 1, z: 1}:  false,
		{x: 1, y: 3, z: 1}:  false,
		{x: 2, y: 3, z: 1}:  false,
		{x: 3, y: 3, z: 1}:  false,
	}

	actual = data.AllAdjacentVertices()
	// compare by element because the test output is hard to read otherwise
	// loop through actual to confirm actual is a subset of expected
	for k, v := range actual {
		v2, ok := expected[k]
		assert.Truef(t, ok, "Key %v found in actual but not in expected", k)
		assert.Equalf(t, v2, v, "Key %v has value %t in actual and %t in expected", k, v, v2)
	}

	// loop through actual to confirm expected is a subset of actual
	// after this loop, this confirms that the maps are equal
	for k, v := range expected {
		v2, ok := actual[k]
		assert.Truef(t, ok, "Key %v found in expected but not in actual", k)
		assert.Equalf(t, v2, v, "Key %v has value %t in expected and %t in actual", k, v, v2)
	}

}

func Test_PuzzleInput_NextGeneration(t *testing.T) {
	var data PuzzleInput
	var expected PuzzleInput
	var actual PuzzleInput
	var err error

	data, err = Parse("../../inputs/day17.example.txt")
	assert.NoError(t, err)

	expected = PuzzleInput{
		Cubes: map[Vertex]bool{
			{x: 0, y: 1, z: -1}: true,
			{x: 2, y: 2, z: -1}: true,
			{x: 1, y: 3, z: -1}: true,

			{x: 0, y: 1, z: 0}: true,
			{x: 2, y: 1, z: 0}: true,
			{x: 1, y: 2, z: 0}: true,
			{x: 2, y: 2, z: 0}: true,
			{x: 1, y: 3, z: 0}: true,

			{x: 0, y: 1, z: 1}: true,
			{x: 2, y: 2, z: 1}: true,
			{x: 1, y: 3, z: 1}: true,
		},
	}

	actual = data.NextGeneration()

	// compare by element because the test output is hard to read otherwise
	// loop through actual to confirm actual is a subset of expected
	for k, v := range actual.Cubes {
		v2, ok := expected.Cubes[k]
		assert.Truef(t, ok, "Key %v found in actual but not in expected", k)
		assert.Equalf(t, v2, v, "Key %v has value %t in actual and %t in expected", k, v, v2)
	}

	// loop through actual to confirm expected is a subset of actual
	// after this loop, this confirms that the maps are equal
	for k, v := range expected.Cubes {
		v2, ok := actual.Cubes[k]
		assert.Truef(t, ok, "Key %v found in expected but not in actual", k)
		assert.Equalf(t, v2, v, "Key %v has value %t in expected and %t in actual", k, v, v2)
	}
}
