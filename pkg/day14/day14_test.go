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
package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseMask(t *testing.T) {
	test := "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	expected := mask{
		ones:   []int{6},
		zeros:  []int{1},
		floats: []int{35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 5, 4, 3, 2, 0},
	}

	actual, err := parseMask(test)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_ParseMemset(t *testing.T) {
	test := "mem[8] = 11"
	expected := memset{addr: 8, value: 11}

	actual, err := parseMemset(test)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Parse(t *testing.T) {
	actual, err := Parse("../../inputs/day14.example01.txt")
	expected := program{
		mask:   nil,
		memory: make(map[int]int),
		instructions: []instruction{
			&mask{
				ones:   []int{6},
				zeros:  []int{1},
				floats: []int{35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 5, 4, 3, 2, 0},
			},
			&memset{
				addr:  8,
				value: 11,
			},
			&memset{
				addr:  7,
				value: 101,
			},
			&memset{
				addr:  8,
				value: 0,
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Solution1(t *testing.T) {
	prog, err := Parse("../../inputs/day14.example01.txt")
	assert.NoError(t, err)

	actual, err := Solution1(prog)
	assert.NoError(t, err)

	expected := "165"
	assert.Equal(t, expected, actual)
}
