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
package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseFile(t *testing.T) {
	testPath := "../../inputs/day09.example.txt"
	expected := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	data, err := ParseIntFile(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)

	testPath = "../../inputs/day10.example01.txt"
	expected = []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

	data, err = ParseIntFile(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}
