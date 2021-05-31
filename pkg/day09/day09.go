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
package day09

import (
	"errors"
	"strconv"
)

func Solution1(xmas []int, lookback int) (string, error) {
	valid, _, num := processXmas(xmas, lookback)
	if valid {
		return "", errors.New("no solution")
	}

	return strconv.Itoa(num), nil
}

func Solution2(xmas []int, lookback int) (string, error) {
	valid, index, num := processXmas(xmas, lookback)
	if valid {
		return "", errors.New("no solution")
	}

	//found the number, now to find the contigous set
	for i := 0; i < index-1; i++ {
		sum := xmas[i]
		for j := i + 1; j < index; j++ {
			sum += xmas[j]
			if sum == num {
				small := smallestInRange(xmas, i, j)
				large := largestInRange(xmas, i, j)
				return strconv.Itoa(small + large), nil
			} else if sum > num {
				break
			}
		}
	}

	return "", errors.New("no solution")
}
