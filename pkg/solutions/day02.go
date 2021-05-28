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
	"strings"
)

type Policy struct {
	Num1 int
	Num2 int
	Char rune
}

type Password struct {
	Policy   Policy
	Contents string
}

func checkCharRangePolicy(password Password) bool {
	count := 0
	chars := []rune(password.Contents)
	for _, c := range chars {
		if c == password.Policy.Char {
			count++
		}
	}

	return count >= password.Policy.Num1 && count <= password.Policy.Num2
}

func checkCharPositionPolicy(password Password) bool {
	chars := []rune(password.Contents)
	if len(chars) < password.Policy.Num1 || len(chars) < password.Policy.Num2 {
		return false
	}

	count := 0
	if chars[password.Policy.Num1-1] != password.Policy.Char {
		count++
	}

	if chars[password.Policy.Num2-1] != password.Policy.Char {
		count++
	}

	return count == 1
}

func ParsePassword(line string) (*Password, error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return nil, fmt.Errorf("expected 3 fields to be found in policy but found %d", len(fields))
	}

	//split the first field into the min/max range numbers
	rangeNums := strings.Split(fields[0], "-")
	if len(rangeNums) != 2 {
		return nil, fmt.Errorf("expected only min and max values for policy, but found %d", len(rangeNums))
	}

	num1, err := strconv.Atoi(rangeNums[0])
	if err != nil {
		return nil, err
	}

	num2, err := strconv.Atoi(rangeNums[1])
	if err != nil {
		return nil, err
	}

	//remove the colon after the character, convert to runes
	policyChar := []rune(strings.TrimSuffix(fields[1], ":"))
	if len(policyChar) != 1 {
		return nil, fmt.Errorf("expected single character in policy but found %d", len(policyChar))
	}

	password := &Password{
		Contents: fields[2],
		Policy: Policy{
			Num1: num1,
			Num2: num2,
			Char: policyChar[0],
		},
	}

	return password, nil
}

func Day02Parse(path string) ([]Password, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []Password

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password, err := ParsePassword(scanner.Text())
		if err != nil {
			return nil, err
		}

		data = append(data, *password)
	}

	return data, nil
}

func Day02Solution01(passwords []Password) string {
	count := 0
	for _, password := range passwords {
		if checkCharRangePolicy(password) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func Day02Solution02(passwords []Password) string {
	count := 0
	for _, password := range passwords {
		if checkCharPositionPolicy(password) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func Day02Solutions(path *string) {
	passwords, err := Day02Parse(*path)
	if err != nil {
		log.Fatal(err)
	}

	soln01 := Day02Solution01(passwords)
	fmt.Printf("Solution 1: %s\n", soln01)

	soln02 := Day02Solution02(passwords)
	fmt.Printf("Solution 2: %s\n", soln02)
}
