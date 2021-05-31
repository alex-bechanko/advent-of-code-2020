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
package day02

import (
	"bufio"
	"os"
	"strconv"
)

func ParseFile(path string) ([]password, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []password

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password, err := parsePassword(scanner.Text())
		if err != nil {
			return nil, err
		}

		data = append(data, password)
	}

	return data, nil
}

func Solution1(passwords []password) (string, error) {
	count := 0
	for _, password := range passwords {
		if checkCharRangePolicy(password) {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

func Solution2(passwords []password) (string, error) {
	count := 0
	for _, password := range passwords {
		if checkCharPositionPolicy(password) {
			count++
		}
	}

	return strconv.Itoa(count), nil
}
