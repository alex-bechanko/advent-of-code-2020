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
package day11

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_raytraceSeat(t *testing.T) {
	data := `.......#.
			 ...#.....
	   		 .#.......
	  		 .........
	 		 ..#L....#
			 ....#....
			 .........
			 #........
			 ...#.....`
	seats, err := parseReader(strings.NewReader(data))
	assert.NoError(t, err)

	x := 3
	y := 4
	var tile string

	assert.Equal(t, emptySeat, seats.layout[y][x])

	tile = seats.raytraceSeat(x, y, -1, -1) // up-left
	assert.Equal(t, occupiedSeat, tile)

	tile = seats.raytraceSeat(x, y, 0, -1) //up
	assert.Equal(t, occupiedSeat, tile)

	tile = seats.raytraceSeat(x, y, +1, -1) //up-right
	assert.Equal(t, occupiedSeat, tile)

	tile = seats.raytraceSeat(x, y, +1, 0) // right
	assert.Equal(t, occupiedSeat, tile)

	tile = seats.raytraceSeat(x, y, +1, +1) // down-right
	assert.Equal(t, occupiedSeat, tile)

	tile = seats.raytraceSeat(x, y, 0, +1) // down
	assert.Equal(t, occupiedSeat, tile)

	tile = seats.raytraceSeat(x, y, -1, +1) // down-left
	assert.Equal(t, occupiedSeat, tile)

	tile = seats.raytraceSeat(x, y, -1, 0) // left
	assert.Equal(t, occupiedSeat, tile)
}

func Test_numOccupiedSeatsRaytrace(t *testing.T) {
	var data string
	var seats seatLayout
	var err error
	var numAdj int

	data = `.......#.
			...#.....
			.#.......
			.........
			..#L....#
			....#....
			.........
			#........
			...#.....`
	seats, err = parseReader(strings.NewReader(data))
	assert.NoError(t, err)

	numAdj, err = seats.numOccupiedSeatsRaytrace(3, 4)
	assert.NoError(t, err)

	assert.Equal(t, 8, numAdj)

	data = `.............
			.L.L.#.#.#.#.
			.............`
	seats, err = parseReader(strings.NewReader(data))
	assert.NoError(t, err)

	numAdj, err = seats.numOccupiedSeatsRaytrace(1, 1)
	assert.NoError(t, err)

	assert.Equal(t, 0, numAdj)

	data = `#.##.##.##
			#######.##
			#.#.#..#..
			####.##.##
			#.##.##.##
			#.#####.##
			..#.#.....
			##########
			#.######.#
			#.#####.##`
	seats, err = parseReader(strings.NewReader(data))
	assert.NoError(t, err)

	numAdj, err = seats.numOccupiedSeatsRaytrace(2, 0)
	assert.NoError(t, err)

	assert.Equal(t, 5, numAdj)

}

func Test_Solution1(t *testing.T) {
	data := `L.LL.LL.LL
			LLLLLLL.LL
			L.L.L..L..
			LLLL.LL.LL
			L.LL.LL.LL
			L.LLLLL.LL
			..L.L.....
			LLLLLLLLLL
			L.LLLLLL.L
			L.LLLLL.LL`
	seats, err := parseReader(strings.NewReader(data))
	assert.NoError(t, err)

	soln, err := Solution1(seats)
	assert.NoError(t, err)
	assert.Equal(t, "37", soln)
}

func Test_Solution2(t *testing.T) {
	data := `L.LL.LL.LL
			LLLLLLL.LL
			L.L.L..L..
			LLLL.LL.LL
			L.LL.LL.LL
			L.LLLLL.LL
			..L.L.....
			LLLLLLLLLL
			L.LLLLLL.L
			L.LLLLL.LL`
	seats, err := parseReader(strings.NewReader(data))
	assert.NoError(t, err)

	soln, err := Solution2(seats)
	assert.NoError(t, err)
	assert.Equal(t, "26", soln)
}
