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
package day06

import (
	"bufio"
	"os"
	"strconv"
)

func ParseFile(path string) ([]planeGroup, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	groups := make([]planeGroup, 0)
	group := newPlaneGroup()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			groups = append(groups, group)
			group = newPlaneGroup()
			continue
		}

		group.addPassenger(newPassenger(line))
	}
	groups = append(groups, group)
	return groups, nil
}

func Solution1(plane []planeGroup) (string, error) {
	total := 0
	for _, grp := range plane {
		total += len(grp.anyoneAnswers)
	}
	return strconv.Itoa(total), nil
}

func Solution2(plane []planeGroup) (string, error) {
	total := 0
	for _, grp := range plane {
		total += len(grp.everyoneAnswers)
	}
	return strconv.Itoa(total), nil
}
