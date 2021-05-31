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
package common

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func ParseIntReader(r io.Reader) ([]int, error) {
	data := make([]int, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		data = append(data, num)
	}

	return data, nil
}

func ParseIntFile(path string) ([]int, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ParseIntReader(file)
}
