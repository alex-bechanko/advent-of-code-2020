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
	"strconv"
	"strings"
)

func passportYr(passport map[string]string, field string) (int, bool) {
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

func validByr(passport map[string]string) bool {
	yr, ok := passportYr(passport, "byr")
	if !ok {
		return false
	}

	return yr >= 1920 && yr <= 2002
}

func validIyr(passport map[string]string) bool {
	yr, ok := passportYr(passport, "iyr")
	if !ok {
		return false
	}

	return yr >= 2010 && yr <= 2020
}

func validEyr(passport map[string]string) bool {
	yr, ok := passportYr(passport, "eyr")
	if !ok {
		return false
	}

	return yr >= 2020 && yr <= 2030
}

func validHgt(passport map[string]string) bool {
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

func validHcl(passport map[string]string) bool {
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

func validEcl(passport map[string]string) bool {
	val, ok := passport["ecl"]
	if !ok {
		return false
	}

	return val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth"
}

func validPid(passport map[string]string) bool {
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
