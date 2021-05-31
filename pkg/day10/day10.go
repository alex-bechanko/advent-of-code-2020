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
package day10

import (
	"errors"
	"sort"
	"strconv"
)

func Solution1(adapters []int) (string, error) {

	// calculate the built-in adapter joltage, then add it
	builtInAdapter := 0
	for _, jolt := range adapters {
		if jolt > builtInAdapter {
			builtInAdapter = jolt
		}
	}

	builtInAdapter += 3
	adapters = append(adapters, builtInAdapter)

	//add the 0 joltage outlet as well
	adapters = append(adapters, 0)

	// problem requires all adapters are used, therefore the order will be a sorted list of the adapters
	sort.Ints(adapters)

	joltDifferenceCounter := map[int]int{}
	for i := 0; i < len(adapters)-1; i++ {
		diff := adapters[i+1] - adapters[i]
		n, ok := joltDifferenceCounter[diff]
		if !ok {
			joltDifferenceCounter[diff] = 1
		} else {
			joltDifferenceCounter[diff] = n + 1
		}
	}

	one, ok := joltDifferenceCounter[1]
	if !ok {
		return "", errors.New("No adapters have a joltage difference of one")
	}
	three, ok := joltDifferenceCounter[3]
	if !ok {
		return "", errors.New("No adapters have a joltage difference of three")
	}

	return strconv.Itoa(one * three), nil
}

func Solution2(adapters []int) (string, error) {
	//add the 0 joltage outlet
	adapters = append(adapters, 0)

	// calculate the built-in adapter joltage, then add it to the path count
	builtInAdapter := 0
	for _, jolt := range adapters {
		if jolt > builtInAdapter {
			builtInAdapter = jolt
		}
	}

	builtInAdapter += 3

	pathCount := make(map[int]int)
	pathCount[builtInAdapter] = 1

	// adapters that can be used as inputs should be on the lower end, to guarantee that
	// path count has already counted the paths of the adapters it could be inputs for
	sort.Ints(adapters)

	for i := len(adapters) - 1; i >= 0; i-- {
		numPaths := 0
		inp := adapters[i]

		var recv int

		recv = pathCount[inp+1]
		numPaths += recv

		recv = pathCount[inp+2]
		numPaths += recv

		recv = pathCount[inp+3]
		numPaths += recv

		pathCount[inp] = numPaths
	}

	return strconv.Itoa(pathCount[0]), nil
}
