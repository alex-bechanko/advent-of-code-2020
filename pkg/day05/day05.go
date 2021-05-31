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
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseFile(path string) ([]seat, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var seats []seat

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 10 {
			return nil, fmt.Errorf("found input line that was not 10 characters in length: %s", line)
		}

		var rs []intComparison
		var cs []intComparison
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

		seat, err := newSeat(rs, cs)
		if err != nil {
			return nil, err
		}

		seats = append(seats, seat)
	}

	return seats, nil
}

func Solution1(seats []seat) (string, error) {
	max := 0

	for _, seat := range seats {
		if seat.id > max {
			max = seat.id
		}
	}

	return strconv.Itoa(max), nil
}

func Solution2(seats []seat) (string, error) {
	ids := make([]int, len(seats))
	for i, seat := range seats {
		ids[i] = seat.id
	}

	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		leftId := ids[i]
		rightId := ids[i+1]
		if rightId-leftId == 2 {
			return strconv.Itoa(leftId + 1), nil
		}
	}

	return "", fmt.Errorf("solution not found")

}
