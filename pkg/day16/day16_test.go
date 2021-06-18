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
package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseRange(t *testing.T) {

	var data string
	var expected RuleRange
	var actual RuleRange
	var err error

	data = "1-3"
	expected = RuleRange{Low: 1, High: 3}

	actual, err = parseRange(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

	data = "5-7"
	expected = RuleRange{Low: 5, High: 7}

	actual, err = parseRange(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_parseRule(t *testing.T) {
	var data string
	var expected Rule
	var actual Rule
	var err error

	data = "class: 1-3 or 5-7\n"
	expected = Rule{
		Name:   "class",
		Range1: RuleRange{Low: 1, High: 3},
		Range2: RuleRange{Low: 5, High: 7},
	}
	actual, err = parseRule(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

	data = "row: 6-11 or 33-44\n"
	expected = Rule{
		Name:   "row",
		Range1: RuleRange{Low: 6, High: 11},
		Range2: RuleRange{Low: 33, High: 44},
	}
	actual, err = parseRule(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

	data = "seat: 13-40 or 45-50\n"
	expected = Rule{
		Name:   "seat",
		Range1: RuleRange{Low: 13, High: 40},
		Range2: RuleRange{Low: 45, High: 50},
	}
	actual, err = parseRule(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_ParseCommaDelimitedInts(t *testing.T) {
	var data string
	var expected []int
	var actual []int
	var err error

	data = "7,1,14"
	expected = []int{7, 1, 14}
	actual, err = ParseCommaDelimitedInts(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

	data = "  7,  3,   47 "
	expected = []int{7, 3, 47}
	actual, err = ParseCommaDelimitedInts(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Parse(t *testing.T) {
	var data string
	var expected PuzzleInput
	var actual PuzzleInput
	var err error

	data = "../../inputs/day16.example01.txt"
	expected = PuzzleInput{
		MyTicket: []int{7, 1, 14},
		NearbyTickets: [][]int{
			{7, 3, 47},
			{40, 4, 50},
			{55, 2, 20},
			{38, 6, 12},
		},
		Rules: map[string]Rule{
			"class": {
				Name: "class",
				Range1: RuleRange{
					Low:  1,
					High: 3,
				},
				Range2: RuleRange{
					Low:  5,
					High: 7,
				},
			},
			"row": {
				Name: "row",
				Range1: RuleRange{
					Low:  6,
					High: 11,
				},
				Range2: RuleRange{
					Low:  33,
					High: 44,
				},
			},
			"seat": {
				Name: "seat",
				Range1: RuleRange{
					Low:  13,
					High: 40,
				},
				Range2: RuleRange{
					Low:  45,
					High: 50,
				},
			},
		},
	}

	actual, err = Parse(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Solution1(t *testing.T) {
	var filepath string
	var data PuzzleInput
	var expected string
	var actual string
	var err error

	filepath = "../../inputs/day16.example01.txt"
	expected = "71"

	data, err = Parse(filepath)
	assert.NoError(t, err)

	actual, err = Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

}

func Test_Solution2(t *testing.T) {
	var filepath string
	var data PuzzleInput
	var expected string
	var actual string
	var err error

	filepath = "../../inputs/day16.txt"
	expected = "2564529489989"

	data, err = Parse(filepath)
	assert.NoError(t, err)

	actual, err = Solution2(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

}
