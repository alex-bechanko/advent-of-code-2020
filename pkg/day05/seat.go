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
package day05

import (
	"fmt"
	"math"
)

const Front string = "F"
const Back string = "B"
const Left string = "L"
const Right string = "R"

type intComparison int

const (
	High intComparison = iota
	Low
)

type seat struct {
	rows   []intComparison
	cols   []intComparison
	row    int
	column int
	id     int
}

func newSeat(rows, cols []intComparison) (seat, error) {

	row, err := findNum(0, 127, rows)
	if err != nil {
		return seat{}, err
	}

	col, err := findNum(0, 7, cols)
	if err != nil {
		return seat{}, err
	}

	return seat{
		rows:   rows,
		cols:   cols,
		row:    row,
		column: col,
		id:     row*8 + col,
	}, nil
}

func findNum(min, max int, directions []intComparison) (int, error) {
	numRange := max - min + 1
	if numRange&(numRange-1) != 0 {
		return 0, fmt.Errorf("numbers between min and max is not a power of 2, can't find exact number. min: %d, max: %d", min, max)
	}

	power := int(math.Log2(float64(numRange)))
	if len(directions) != power {
		return 0, fmt.Errorf("number of directions provided is not enough to find exact number. min: %d, max: %d, numDirections: %d", min, max, len(directions))
	}

	for i := 0; i < len(directions)-1; i++ {
		direction := directions[i]

		if direction == Low {
			max = max - (max-min+1)/2
		} else {
			min = min + (max-min+1)/2
		}
	}

	if directions[len(directions)-1] == Low {
		return min, nil
	}

	return max, nil
}
