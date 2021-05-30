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
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day04Parse(t *testing.T) {
	testPath := "../../inputs/day04.example.txt"
	expected := []map[string]string{
		{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"},
		{"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"},
		{"hcl": "#ae17e1", "iyr": "2013", "eyr": "2024", "ecl": "brn", "pid": "760753108", "byr": "1931", "hgt": "179cm"},
		{"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648", "iyr": "2011", "ecl": "brn", "hgt": "59in"},
	}
	data, err := Day04Parse(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_Day04Solution01(t *testing.T) {
	data := []map[string]string{
		{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"},
		{"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"},
		{"hcl": "#ae17e1", "iyr": "2013", "eyr": "2024", "ecl": "brn", "pid": "760753108", "byr": "1931", "hgt": "179cm"},
		{"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648", "iyr": "2011", "ecl": "brn", "hgt": "59in"},
	}
	expected := "2"
	actual, err := Day04Solution01(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Day04Solution02(t *testing.T) {
	data := []map[string]string{
		{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"},
		{"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"},
		{"hcl": "#ae17e1", "iyr": "2013", "eyr": "2024", "ecl": "brn", "pid": "760753108", "byr": "1931", "hgt": "179cm"},
		{"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648", "iyr": "2011", "ecl": "brn", "hgt": "59in"},
	}
	expected := "2"
	actual, err := Day04Solution02(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
