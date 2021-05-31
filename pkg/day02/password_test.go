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

func Test_checkCharRangePolicy(t *testing.T) {
	data := password{
		contents: "abcde",
		policy:   policy{num1: 1, num2: 3, char: 'a'},
	}

	assert.True(t, checkCharRangePolicy(data))
}
func Test_checkCharPositionPolicy(t *testing.T) {
	data := password{
		contents: "abcde",
		policy:   policy{num1: 1, num2: 3, char: 'a'},
	}

	assert.True(t, checkCharPositionPolicy(data))
}
func Test_parsePassword(t *testing.T) {
	expected := password{
		contents: "abcde",
		policy:   policy{num1: 1, num2: 3, char: 'a'},
	}
	data, err := parsePassword("1-3 a: abcde")

	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}
