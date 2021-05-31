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
package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseFile(t *testing.T) {
	testPath := "../../inputs/day02.example.txt"
	expected := []password{
		{
			contents: "abcde",
			policy:   policy{num1: 1, num2: 3, char: 'a'},
		},
		{
			contents: "cdefg",
			policy:   policy{num1: 1, num2: 3, char: 'b'},
		},
		{
			contents: "ccccccccc",
			policy:   policy{num1: 2, num2: 9, char: 'c'},
		},
	}
	data, err := ParseFile(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_Solution1(t *testing.T) {
	data := []password{
		{
			contents: "abcde",
			policy:   policy{num1: 1, num2: 3, char: 'a'},
		},
		{
			contents: "cdefg",
			policy:   policy{num1: 1, num2: 3, char: 'b'},
		},
		{
			contents: "ccccccccc",
			policy:   policy{num1: 2, num2: 9, char: 'c'},
		},
	}
	expected := "2"
	actual, err := Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Solution2(t *testing.T) {
	data := []password{
		{
			contents: "abcde",
			policy:   policy{num1: 1, num2: 3, char: 'a'},
		},
		{
			contents: "cdefg",
			policy:   policy{num1: 1, num2: 3, char: 'b'},
		},
		{
			contents: "ccccccccc",
			policy:   policy{num1: 2, num2: 9, char: 'c'},
		},
	}
	expected := "1"
	actual, err := Solution2(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
