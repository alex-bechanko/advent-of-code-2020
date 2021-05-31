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
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseFile(path string) (forest, error) {

	forest := forest{}

	file, err := os.Open(path)
	if err != nil {
		return forest, err
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

	forest.rows = rows
	forest.rowLength = rowlength
	forest.columnHeight = columnheight

	return forest, nil

}

func Solution1(forest forest) (string, error) {
	return strconv.Itoa(traverseSlope(forest, 3, 1)), nil
}

func Solution2(forest forest) (string, error) {
	answer := 1
	answer *= traverseSlope(forest, 1, 1)
	answer *= traverseSlope(forest, 3, 1)
	answer *= traverseSlope(forest, 5, 1)
	answer *= traverseSlope(forest, 7, 1)
	answer *= traverseSlope(forest, 1, 2)

	return strconv.Itoa(answer), nil
}
