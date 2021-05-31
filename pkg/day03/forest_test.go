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
package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_traverseSlope(t *testing.T) {
	data := forest{
		rowLength:    11,
		columnHeight: 11,
		rows: [][]string{
			{".", ".", "#", "#", ".", ".", ".", ".", ".", ".", "."},
			{"#", ".", ".", ".", "#", ".", ".", ".", "#", ".", "."},
			{".", "#", ".", ".", ".", ".", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "#"},
			{".", "#", ".", ".", ".", "#", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", "#", ".", ".", ".", ".", "."},
			{".", "#", ".", "#", ".", "#", ".", ".", ".", ".", "#"},
			{".", "#", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
			{"#", ".", "#", "#", ".", ".", ".", "#", ".", ".", "."},
			{"#", ".", ".", ".", "#", "#", ".", ".", ".", ".", "#"},
			{".", "#", ".", ".", "#", ".", ".", ".", "#", ".", "#"},
		},
	}
	expected := 7
	actual := traverseSlope(data, 3, 1)
	assert.Equal(t, expected, actual)
}

func Test_isTree(t *testing.T) {
	data := forest{
		rowLength:    11,
		columnHeight: 11,
		rows: [][]string{
			{".", ".", "#", "#", ".", ".", ".", ".", ".", ".", "."},
			{"#", ".", ".", ".", "#", ".", ".", ".", "#", ".", "."},
			{".", "#", ".", ".", ".", ".", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "#"},
			{".", "#", ".", ".", ".", "#", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", "#", ".", ".", ".", ".", "."},
			{".", "#", ".", "#", ".", "#", ".", ".", ".", ".", "#"},
			{".", "#", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
			{"#", ".", "#", "#", ".", ".", ".", "#", ".", ".", "."},
			{"#", ".", ".", ".", "#", "#", ".", ".", ".", ".", "#"},
			{".", "#", ".", ".", "#", ".", ".", ".", "#", ".", "#"},
		},
	}
	assert.True(t, data.isTree(2, 0))
	assert.True(t, data.isTree(0, 1))
	assert.False(t, data.isTree(0, 0))
	assert.False(t, data.isTree(0, 10))
}
