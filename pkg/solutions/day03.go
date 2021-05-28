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
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Forest struct {
	Rows         [][]string
	RowLength    int
	ColumnHeight int
}

func (f Forest) String() string {
	var sb strings.Builder
	for _, r := range f.Rows {
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

func (f Forest) IsTree(x, y int) bool {

	// forests are allowed to repeat on the x axis
	x = x % f.RowLength
	if x < 0 {
		x += f.RowLength
	}

	// forests don't repeat on the y-axis, panic if it occurs
	if y < 0 || y >= f.ColumnHeight {
		log.Fatalf("Attempted to use an out-of-bounds y coordinate for forests: %d", y)
	}

	return f.Rows[y][x] == "#"

}

func Day03Parse(path string) (*Forest, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows [][]string
	rowlength := 0
	columnheight := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		columnheight++

		line := scanner.Text()
		rowlength = len(line)

		rows = append(rows, strings.Split(line, ""))
	}

	forest := &Forest{
		Rows:         rows,
		RowLength:    rowlength,
		ColumnHeight: columnheight,
	}

	return forest, nil

}

func TraverseSlope(forest *Forest, dx, dy int) int {
	x := 0
	y := 0

	trees := 0
	for y < forest.ColumnHeight {
		if forest.IsTree(x, y) {
			trees++
		}
		x += dx
		y += dy
	}

	return trees
}

func Day03Solution01(forest *Forest) string {
	return strconv.Itoa(TraverseSlope(forest, 3, 1))
}

func Day03Solution02(forest *Forest) string {
	answer := 1
	answer *= TraverseSlope(forest, 1, 1)
	answer *= TraverseSlope(forest, 3, 1)
	answer *= TraverseSlope(forest, 5, 1)
	answer *= TraverseSlope(forest, 7, 1)
	answer *= TraverseSlope(forest, 1, 2)

	return strconv.Itoa(answer)
}

func Day03Solutions(path *string) {
	forest, err := Day03Parse(*path)
	if err != nil {
		log.Fatal(err)
	}

	soln01 := Day03Solution01(forest)
	fmt.Printf("Solution 1: %s\n", soln01)

	soln02 := Day03Solution02(forest)
	fmt.Printf("Solution 1: %s\n", soln02)

}
