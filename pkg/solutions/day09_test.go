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

func Test_Day09Parse(t *testing.T) {
	testPath := "../../inputs/day09.example.txt"
	expected := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	data, err := Day09Parse(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_Day09Solution01(t *testing.T) {
	data := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	lookback := 5

	expected := "127"
	actual, err := Day09Solution01(data, lookback)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Day09Solution02(t *testing.T) {
	data := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	lookback := 5

	expected := "62"
	actual, err := Day09Solution02(data, lookback)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
