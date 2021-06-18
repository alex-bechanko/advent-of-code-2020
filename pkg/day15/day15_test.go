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
package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	testPath := "../../inputs/day15.example.txt"
	expected := MemoryGame{
		LastSpoken:      0,
		Spoken:          make(map[int]int),
		StartingNumbers: []int{0, 3, 6},
	}

	data, err := Parse(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_PlayGame(t *testing.T) {
	testPath := "../../inputs/day15.example.txt"
	data, err := Parse(testPath)
	assert.NoError(t, err)

	answer, err := PlayGame(data, 10)
	assert.NoError(t, err)

	assert.Equal(t, "0", answer)
}
