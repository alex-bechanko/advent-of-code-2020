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
package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day02Parse(t *testing.T) {
	testPath := "../../inputs/day02.example.txt"
	expected := []Password{
		{
			Contents: "abcde",
			Policy:   Policy{Num1: 1, Num2: 3, Char: 'a'},
		},
		{
			Contents: "cdefg",
			Policy:   Policy{Num1: 1, Num2: 3, Char: 'b'},
		},
		{
			Contents: "ccccccccc",
			Policy:   Policy{Num1: 2, Num2: 9, Char: 'c'},
		},
	}
	data, err := Day02Parse(testPath)
	assert.NoErrorf(t, err, "Error occurred parsing %s", testPath)
	assert.Equal(t, expected, data)
}

func Test_Day02Solution01(t *testing.T) {
	data := []Password{
		{
			Contents: "abcde",
			Policy:   Policy{Num1: 1, Num2: 3, Char: 'a'},
		},
		{
			Contents: "cdefg",
			Policy:   Policy{Num1: 1, Num2: 3, Char: 'b'},
		},
		{
			Contents: "ccccccccc",
			Policy:   Policy{Num1: 2, Num2: 9, Char: 'c'},
		},
	}
	expected := "2"
	actual := Day02Solution01(data)

	assert.Equal(t, expected, actual)
}

func Test_Day02Solution02(t *testing.T) {
	data := []Password{
		{
			Contents: "abcde",
			Policy:   Policy{Num1: 1, Num2: 3, Char: 'a'},
		},
		{
			Contents: "cdefg",
			Policy:   Policy{Num1: 1, Num2: 3, Char: 'b'},
		},
		{
			Contents: "ccccccccc",
			Policy:   Policy{Num1: 2, Num2: 9, Char: 'c'},
		},
	}
	expected := "1"
	actual := Day02Solution02(data)

	assert.Equal(t, expected, actual)
}
