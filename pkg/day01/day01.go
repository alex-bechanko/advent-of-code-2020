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
package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ParseFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		data = append(data, num)
	}

	return data, nil
}

func Solution1(data []int) (string, error) {
	for i := 0; i+1 < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == 2020 {
				return strconv.Itoa(data[i] * data[j]), nil
			}
		}
	}
	return "", fmt.Errorf("No solution found for part1")
}

func Solution2(data []int) (string, error) {
	for i := 0; i+2 < len(data); i++ {
		for j := i + 1; j+1 < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				if data[i]+data[j]+data[k] == 2020 {
					return strconv.Itoa(data[i] * data[j] * data[k]), nil
				}
			}
		}
	}
	return "", fmt.Errorf("No solution found for part1")
}
