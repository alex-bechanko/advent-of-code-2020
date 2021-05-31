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
package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseFile(t *testing.T) {
	testPath := "../../inputs/day06.example.txt"
	expected := []PlaneGroup{
		{
			AnyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
			EveryoneAnswers: map[string]bool{"a": true, "b": true, "c": true},
			Count:           1,
		},
		{
			AnyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
			EveryoneAnswers: map[string]bool{},
			Count:           3,
		},
		{
			AnyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
			EveryoneAnswers: map[string]bool{"a": true},
			Count:           2,
		},
		{
			AnyoneAnswers:   map[string]bool{"a": true},
			EveryoneAnswers: map[string]bool{"a": true},
			Count:           4,
		},
		{
			AnyoneAnswers:   map[string]bool{"b": true},
			EveryoneAnswers: map[string]bool{"b": true},
			Count:           1,
		},
	}

	data, err := ParseFile(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_Solution1(t *testing.T) {
	data := []PlaneGroup{
		{
			AnyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
			EveryoneAnswers: map[string]bool{"a": true, "b": true, "c": true},
			Count:           1,
		},
		{
			AnyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
			EveryoneAnswers: map[string]bool{},
			Count:           3,
		},
		{
			AnyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
			EveryoneAnswers: map[string]bool{"a": true},
			Count:           2,
		},
		{
			AnyoneAnswers:   map[string]bool{"a": true},
			EveryoneAnswers: map[string]bool{"a": true},
			Count:           4,
		},
		{
			AnyoneAnswers:   map[string]bool{"b": true},
			EveryoneAnswers: map[string]bool{"b": true},
			Count:           1,
		},
	}
	expected := "11"
	actual, err := Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Solution2(t *testing.T) {
	data := []PlaneGroup{
		{
			AnyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
			EveryoneAnswers: map[string]bool{"a": true, "b": true, "c": true},
			Count:           1,
		},
		{
			AnyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
			EveryoneAnswers: map[string]bool{},
			Count:           3,
		},
		{
			AnyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
			EveryoneAnswers: map[string]bool{"a": true},
			Count:           2,
		},
		{
			AnyoneAnswers:   map[string]bool{"a": true},
			EveryoneAnswers: map[string]bool{"a": true},
			Count:           4,
		},
		{
			AnyoneAnswers:   map[string]bool{"b": true},
			EveryoneAnswers: map[string]bool{"b": true},
			Count:           1,
		},
	}
	expected := "6"
	actual, err := Solution2(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
