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

func PassportYr(passport map[string]string, field string) (int, bool) {
	val, ok := passport[field]
	if !ok {
		return 0, false
	}

	yr, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}

	return yr, true
}

func ValidByr(passport map[string]string) bool {
	yr, ok := PassportYr(passport, "byr")
	if !ok {
		return false
	}

	return yr >= 1920 && yr <= 2002
}

func ValidIyr(passport map[string]string) bool {
	yr, ok := PassportYr(passport, "iyr")
	if !ok {
		return false
	}

	return yr >= 2010 && yr <= 2020
}

func ValidEyr(passport map[string]string) bool {
	yr, ok := PassportYr(passport, "eyr")
	if !ok {
		return false
	}

	return yr >= 2020 && yr <= 2030
}

func ValidHgt(passport map[string]string) bool {
	val, ok := passport["hgt"]
	if !ok {
		return false
	}

	if strings.HasSuffix(val, "cm") {
		hgt, err := strconv.Atoi(strings.TrimSuffix(val, "cm"))
		if err != nil {
			return false
		}
		return hgt >= 150 && hgt <= 193

	} else if strings.HasSuffix(val, "in") {
		hgt, err := strconv.Atoi(strings.TrimSuffix(val, "in"))
		if err != nil {
			return false
		}
		return hgt >= 59 && hgt <= 76

	}

	return false
}

func ValidHcl(passport map[string]string) bool {
	hcl, ok := passport["hcl"]
	if !ok || len(hcl) != 7 || !strings.HasPrefix(hcl, "#") {
		return false
	}

	hclChars := strings.Split(strings.TrimPrefix(hcl, "#"), "")
	for _, c := range hclChars {
		if !strings.ContainsAny(c, "abcdef0123456789") {
			return false
		}
	}
	return true
}

func ValidEcl(passport map[string]string) bool {
	val, ok := passport["ecl"]
	if !ok {
		return false
	}

	return val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth"
}

func ValidPid(passport map[string]string) bool {
	val, ok := passport["pid"]
	if !ok || len(val) != 9 {
		return false
	}
	pid := strings.Split(val, "")

	for _, c := range pid {
		if !strings.ContainsAny(c, "0123456789") {
			return false
		}
	}

	return true
}

func Solution2(passports []map[string]string) (string, error) {
	count := 0
	validationFuncs := [](func(map[string]string) bool){ValidByr, ValidEcl, ValidEyr, ValidHcl, ValidHgt, ValidIyr, ValidPid}
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
