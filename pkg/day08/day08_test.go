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
package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseFile(t *testing.T) {
	testPath := "../../inputs/day08.example.txt"
	expected := program{
		ip:  0,
		acc: 0,
		commands: []instruction{
			{op: "nop", arg: "+0"},
			{op: "acc", arg: "+1"},
			{op: "jmp", arg: "+4"},
			{op: "acc", arg: "+3"},
			{op: "jmp", arg: "-3"},
			{op: "acc", arg: "-99"},
			{op: "acc", arg: "+1"},
			{op: "jmp", arg: "-4"},
			{op: "acc", arg: "+6"}},
	}

	data, err := ParseFile(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_Solution1(t *testing.T) {
	data := program{
		ip:  0,
		acc: 0,
		commands: []instruction{
			{op: "nop", arg: "+0"},
			{op: "acc", arg: "+1"},
			{op: "jmp", arg: "+4"},
			{op: "acc", arg: "+3"},
			{op: "jmp", arg: "-3"},
			{op: "acc", arg: "-99"},
			{op: "acc", arg: "+1"},
			{op: "jmp", arg: "-4"},
			{op: "acc", arg: "+6"}},
	}
	expected := "5"
	actual, err := Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Solution2(t *testing.T) {
	data := program{
		ip:  0,
		acc: 0,
		commands: []instruction{
			{op: "nop", arg: "+0"},
			{op: "acc", arg: "+1"},
			{op: "jmp", arg: "+4"},
			{op: "acc", arg: "+3"},
			{op: "jmp", arg: "-3"},
			{op: "acc", arg: "-99"},
			{op: "acc", arg: "+1"},
			{op: "jmp", arg: "-4"},
			{op: "acc", arg: "+6"}},
	}
	expected := "8"
	actual, err := Solution2(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
