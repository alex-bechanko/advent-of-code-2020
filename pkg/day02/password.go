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
	"fmt"
	"strconv"
	"strings"
)

type policy struct {
	num1 int
	num2 int
	char rune
}

type password struct {
	policy   policy
	contents string
}

func checkCharRangePolicy(password password) bool {
	count := 0
	chars := []rune(password.contents)
	for _, c := range chars {
		if c == password.policy.char {
			count++
		}
	}

	return count >= password.policy.num1 && count <= password.policy.num2
}

func checkCharPositionPolicy(password password) bool {
	chars := []rune(password.contents)
	if len(chars) < password.policy.num1 || len(chars) < password.policy.num2 {
		return false
	}

	count := 0
	if chars[password.policy.num1-1] != password.policy.char {
		count++
	}

	if chars[password.policy.num2-1] != password.policy.char {
		count++
	}

	return count == 1
}

func parsePassword(line string) (password, error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return password{}, fmt.Errorf("expected 3 fields to be found in policy but found %d", len(fields))
	}

	//split the first field into the min/max range numbers
	rangeNums := strings.Split(fields[0], "-")
	if len(rangeNums) != 2 {
		return password{}, fmt.Errorf("expected only min and max values for policy, but found %d", len(rangeNums))
	}

	num1, err := strconv.Atoi(rangeNums[0])
	if err != nil {
		return password{}, err
	}

	num2, err := strconv.Atoi(rangeNums[1])
	if err != nil {
		return password{}, err
	}

	//remove the colon after the character, convert to runes
	policyChar := []rune(strings.TrimSuffix(fields[1], ":"))
	if len(policyChar) != 1 {
		return password{}, fmt.Errorf("expected single character in policy but found %d", len(policyChar))
	}

	password := password{
		contents: fields[2],
		policy: policy{
			num1: num1,
			num2: num2,
			char: policyChar[0],
		},
	}

	return password, nil
}
