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

func Test_Day01Parse(t *testing.T) {
	testPath := "../../inputs/day01.example.txt"
	expected := []int{1721, 979, 366, 299, 675, 1456}

	data, err := Day01Parse(testPath)
	assert.NoErrorf(t, err, "Recieved error when parsing %s", testPath, err)

	assert.Equal(t, expected, data, "data does not match expected")
}

func Test_Day01Solution01(t *testing.T) {
	data := []int{1721, 979, 366, 299, 675, 1456}
	expected := "514579"
	actual, err := Day01Solution01(data)
	assert.NoError(t, err, "Error occurred calculating day01 part01 solution")
	assert.Equal(t, expected, actual)

}

func Test_Day01Solution02(t *testing.T) {
	data := []int{1721, 979, 366, 299, 675, 1456}
	expected := "241861950"
	actual, err := Day01Solution02(data)
	assert.NoError(t, err, "Error occurred calculating day01 part02 solution")
	assert.Equal(t, expected, actual)
}
