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
package day03

import (
	"log"
	"strings"
)

type forest struct {
	rows         [][]string
	rowLength    int
	columnHeight int
}

func (f forest) String() string {
	var sb strings.Builder
	for _, r := range f.rows {
		for _, c := range r {
			_, err := sb.WriteString(c)
			if err != nil {
				log.Fatal(err)
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (f forest) isTree(x, y int) bool {

	// forests are allowed to repeat on the x axis
	x = x % f.rowLength
	if x < 0 {
		x += f.rowLength
	}

	// forests don't repeat on the y-axis, panic if it occurs
	if y < 0 || y >= f.columnHeight {
		log.Fatalf("Attempted to use an out-of-bounds y coordinate for forests: %d", y)
	}

	return f.rows[y][x] == "#"

}

func traverseSlope(forest forest, dx, dy int) int {
	x := 0
	y := 0

	trees := 0
	for y < forest.columnHeight {
		if forest.isTree(x, y) {
			trees++
		}
		x += dx
		y += dy
	}

	return trees
}
