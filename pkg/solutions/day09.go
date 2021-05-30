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
	"os"
	"strconv"
)

func Day09Parse(path string) ([]int, error) {
	xmas := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		return xmas, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			return xmas, err
		}

		xmas = append(xmas, num)
	}

	return xmas, nil
}

func SmallestInRange(xmas []int, start, end int) int {
	small := xmas[start]
	for i := start; i <= end; i++ {
		if xmas[i] < small {
			small = xmas[i]
		}
	}

	return small
}

func LargestInRange(xmas []int, start, end int) int {
	large := xmas[start]
	for i := start; i <= end; i++ {
		if xmas[i] > large {
			large = xmas[i]
		}
	}

	return large
}

func ValidEntry(xmas []int, i, lookback int) bool {
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

func ProcessXmas(xmas []int, lookback int) (bool, int, int) {
	if len(xmas) <= lookback {
		log.Fatal("Cannot have a xmas list that is smaller than the lookback")
	}

	for i := lookback; i < len(xmas); i++ {
		if !ValidEntry(xmas, i, lookback) {
			return false, i, xmas[i]
		}
	}

	return true, 0, 0
}

func Day09Solution01(xmas []int, lookback int) string {
	valid, _, num := ProcessXmas(xmas, lookback)
	if valid {
		return "no solution"
	}

	return strconv.Itoa(num)
}

func Day09Solution02(xmas []int, lookback int) string {
	valid, index, num := ProcessXmas(xmas, lookback)
	if valid {
		return "no solution"
	}

	//found the number, now to find the contigous set
	for i := 0; i < index-1; i++ {
		sum := xmas[i]
		for j := i + 1; j < index; j++ {
			sum += xmas[j]
			if sum == num {
				small := SmallestInRange(xmas, i, j)
				large := LargestInRange(xmas, i, j)
				return strconv.Itoa(small + large)
			} else if sum > num {
				break
			}
		}
	}

	return "no solution"
}

func Day09Solutions(path *string, lookback int) {
	xmas, err := Day09Parse(*path)
	if err != nil {
		log.Fatal(err)
	}

	soln01 := Day09Solution01(xmas, lookback)
	fmt.Printf("Solution 1: %s\n", soln01)

	soln02 := Day09Solution02(xmas, lookback)
	fmt.Printf("Solution 2: %s\n", soln02)
}
