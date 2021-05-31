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
package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestInRange(t *testing.T) {
	data := []int{0, 10, 2, 3, -4, 5, 6, 7}
	assert.Equal(t, -4, smallestInRange(data, 0, 7))
}

func Test_largestInRange(t *testing.T) {
	data := []int{0, 10, 2, 3, -4, 5, 6, 7}
	assert.Equal(t, 10, largestInRange(data, 0, 7))
}

func Test_validEntry(t *testing.T) {
	var data []int

	data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.True(t, validEntry(data, 6, 5))

	data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	assert.False(t, validEntry(data, 6, 5))

}

func Test_processXmas(t *testing.T) {
	data := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	lookback := 5

	valid, index, num := processXmas(data, lookback)

	assert.False(t, valid)
	assert.Equal(t, 14, index)
	assert.Equal(t, 127, num)
}
