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
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const Front string = "F"
const Back string = "B"
const Left string = "L"
const Right string = "R"

type IntComparison int

const (
	High IntComparison = iota
	Low
)

type Seat struct {
	Rows   []IntComparison
	Cols   []IntComparison
	Row    int
	Column int
	ID     int
}

func NewSeat(rows, cols []IntComparison) (*Seat, error) {
	row, err := findNum(0, 127, rows)
	if err != nil {
		return nil, err
	}

	col, err := findNum(0, 7, cols)
	if err != nil {
		return nil, err
	}

	return &Seat{
		Rows:   rows,
		Cols:   cols,
		Row:    row,
		Column: col,
		ID:     row*8 + col,
	}, nil
}

func findNum(min, max int, directions []IntComparison) (int, error) {
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

func Day05Parse(path string) (*[]Seat, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var seats []Seat

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 10 {
			return nil, fmt.Errorf("found input line that was not 10 characters in length: %s", line)
		}

		var rs []IntComparison
		var cs []IntComparison
		for i, c := range strings.Split(line, "") {
			if i < 7 {
				if c == Front {
					rs = append(rs, Low)
				} else if c == Back {
					rs = append(rs, High)
				} else {
					return nil, fmt.Errorf("found character or index mismatch. line: %s, char: %s, index: %d", line, c, i)
				}

			} else if i < 10 {
				if c == Left {
					cs = append(cs, Low)
				} else if c == Right {
					cs = append(cs, High)
				}

			} else {
				return nil, fmt.Errorf("found character or index mismatch. line: %s, char: %s, index: %d", line, c, i)
			}
		}

		seat, err := NewSeat(rs, cs)
		if err != nil {
			return nil, err
		}

		seats = append(seats, *seat)
	}

	return &seats, nil
}

func Day05Solution01(seats []Seat) string {
	max := 0

	for _, seat := range seats {
		if seat.ID > max {
			max = seat.ID
		}
	}

	return strconv.Itoa(max)
}

func Day05Solution02(seats []Seat) string {
	ids := make([]int, len(seats))
	for i, seat := range seats {
		ids[i] = seat.ID
	}

	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		leftId := ids[i]
		rightId := ids[i+1]
		if rightId-leftId == 2 {
			return strconv.Itoa(leftId + 1)
		}
	}

	return "solution not found"

}

func Day05Solutions(path *string) {
	seats, err := Day05Parse(*path)
	if err != nil {
		log.Fatal(err)
	}

	soln01 := Day05Solution01(*seats)
	fmt.Printf("Solution 1: %s\n", soln01)

	soln02 := Day05Solution02(*seats)
	fmt.Printf("Solution 2: %s\n", soln02)
}
