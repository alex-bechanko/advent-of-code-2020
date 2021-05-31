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
package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(path string) ([]map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var passports [](map[string]string)
	passport := make(map[string]string)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			passports = append(passports, passport)
			passport = make(map[string]string)

			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			parts := strings.Split(field, ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("field was not parsed into two pieces: %s", field)
			}
			fieldName := parts[0]
			fieldValue := parts[1]

			passport[fieldName] = fieldValue
		}
	}
	passports = append(passports, passport)

	return passports, nil
}

func Solution1(passports []map[string]string) (string, error) {

	count := 0
	expectedFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, passport := range passports {
		valid := true
		for _, f := range expectedFields {
			if _, ok := passport[f]; !ok {
				valid = false
				break
			}
		}

		if valid {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

func Solution2(passports []map[string]string) (string, error) {
	count := 0
	validationFuncs := [](func(map[string]string) bool){validByr, validEcl, validEyr, validHcl, validHgt, validIyr, validPid}
	for _, passport := range passports {
		valid := true
		for _, f := range validationFuncs {
			if !f(passport) {
				valid = false
				break
			}
		}

		if valid {
			count++
		}
	}

	return strconv.Itoa(count), nil
}
