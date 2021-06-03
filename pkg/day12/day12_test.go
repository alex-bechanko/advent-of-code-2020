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
package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	data, err := Parse("../../inputs/day12.example.txt")
	expected := solutionData{
		moves: []move{
			{direction: forward, magnitude: 10},
			{direction: north, magnitude: 3},
			{direction: forward, magnitude: 7},
			{direction: right, magnitude: 90},
			{direction: forward, magnitude: 11},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, data)

}

func Test_Solution1(t *testing.T) {
	data, err := Parse("../../inputs/day12.example.txt")
	expected := "25"
	assert.NoError(t, err)

	soln, err := Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, soln)

}

func Test_Solution2(t *testing.T) {
	data, err := Parse("../../inputs/day12.example.txt")
	expected := "286"
	assert.NoError(t, err)

	soln, err := Solution2(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, soln)
}

func Test_vec2dManhatten(t *testing.T) {
	var v vec2d
	v = vec2d{1, 0}
	assert.Equal(t, float64(1), v.manhatten())

	v = vec2d{1, 1}
	assert.Equal(t, float64(2), v.manhatten())

	v = vec2d{-1, -1}
	assert.Equal(t, float64(2), v.manhatten())
}

func Test_vec2dAdd(t *testing.T) {
	var v vec2d
	var ans vec2d
	v = vec2d{1, 1}
	ans = v.add(1, 4)
	assert.Equal(t, vec2d{2, 5}, ans)
}

func Test_vec2dRotate(t *testing.T) {
	var v vec2d
	var ans vec2d
	delta := 0.000001

	v = vec2d{1, 0}
	ans = v.rotate(180)
	assert.InDelta(t, -1, ans.x, delta)
	assert.InDelta(t, 0, ans.y, delta)

	v = vec2d{-1, 1}
	ans = v.rotate(90)
	assert.InDelta(t, -1, ans.x, delta)
	assert.InDelta(t, -1, ans.y, delta)
}
