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
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_passportYr(t *testing.T) {
	passport := map[string]string{
		"ecl": "gry",
		"pid": "860033327",
		"eyr": "2020",
		"hcl": "#fffffd",
		"byr": "1937",
		"iyr": "2017",
		"cid": "147",
		"hgt": "183cm",
	}
	var yr int
	var valid bool

	yr, valid = passportYr(passport, "eyr")
	assert.True(t, valid)
	assert.Equal(t, 2020, yr)

	yr, valid = passportYr(passport, "byr")
	assert.True(t, valid)
	assert.Equal(t, 1937, yr)

	yr, valid = passportYr(passport, "iyr")
	assert.True(t, valid)
	assert.Equal(t, 2017, yr)

	yr, valid = passportYr(passport, "dontexist")
	assert.False(t, valid)
	assert.Equal(t, 0, yr)
}
func Test_validByr(t *testing.T) {
	passport := map[string]string{
		"ecl": "gry",
		"pid": "860033327",
		"eyr": "2020",
		"hcl": "#fffffd",
		"byr": "1937",
		"iyr": "2017",
		"cid": "147",
		"hgt": "183cm",
	}

	assert.True(t, validByr(passport))

}

func Test_validIyr(t *testing.T) {
	passport := map[string]string{
		"ecl": "gry",
		"pid": "860033327",
		"eyr": "2020",
		"hcl": "#fffffd",
		"byr": "1937",
		"iyr": "2017",
		"cid": "147",
		"hgt": "183cm",
	}

	assert.True(t, validIyr(passport))
}
func Test_validEyr(t *testing.T) {
	passport := map[string]string{
		"ecl": "gry",
		"pid": "860033327",
		"eyr": "2020",
		"hcl": "#fffffd",
		"byr": "1937",
		"iyr": "2017",
		"cid": "147",
		"hgt": "183cm",
	}

	assert.True(t, validEyr(passport))
}
func Test_validHgt(t *testing.T) {
	passport := map[string]string{
		"ecl": "gry",
		"pid": "860033327",
		"eyr": "2020",
		"hcl": "#fffffd",
		"byr": "1937",
		"iyr": "2017",
		"cid": "147",
		"hgt": "183cm",
	}

	assert.True(t, validHgt(passport))
}
func Test_validHcl(t *testing.T) {
	passport := map[string]string{
		"ecl": "gry",
		"pid": "860033327",
		"eyr": "2020",
		"hcl": "#fffffd",
		"byr": "1937",
		"iyr": "2017",
		"cid": "147",
		"hgt": "183cm",
	}

	assert.True(t, validHcl(passport))
}
func Test_validEcl(t *testing.T) {
	passport := map[string]string{
		"ecl": "gry",
		"pid": "860033327",
		"eyr": "2020",
		"hcl": "#fffffd",
		"byr": "1937",
		"iyr": "2017",
		"cid": "147",
		"hgt": "183cm",
	}

	assert.True(t, validEcl(passport))
}
func Test_validPid(t *testing.T) {
	passport := map[string]string{
		"ecl": "gry",
		"pid": "860033327",
		"eyr": "2020",
		"hcl": "#fffffd",
		"byr": "1937",
		"iyr": "2017",
		"cid": "147",
		"hgt": "183cm",
	}

	assert.True(t, validPid(passport))
}
