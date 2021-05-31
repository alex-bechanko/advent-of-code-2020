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

import "log"

func smallestInRange(xmas []int, start, end int) int {
	small := xmas[start]
	for i := start; i <= end; i++ {
		if xmas[i] < small {
			small = xmas[i]
		}
	}

	return small
}

func largestInRange(xmas []int, start, end int) int {
	large := xmas[start]
	for i := start; i <= end; i++ {
		if xmas[i] > large {
			large = xmas[i]
		}
	}

	return large
}

func validEntry(xmas []int, i, lookback int) bool {
	var start int
	if i-lookback < 0 {
		start = 0
	} else {
		start = i - lookback
	}
	for j := start; j < i-1; j++ {
		for k := j + 1; k < i; k++ {
			if xmas[j]+xmas[k] == xmas[i] {
				return true
			}
		}
	}
	return false
}

func processXmas(xmas []int, lookback int) (bool, int, int) {
	if len(xmas) <= lookback {
		log.Fatal("Cannot have a xmas list that is smaller than the lookback")
	}

	for i := lookback; i < len(xmas); i++ {
		if !validEntry(xmas, i, lookback) {
			return false, i, xmas[i]
		}
	}

	return true, 0, 0
}
