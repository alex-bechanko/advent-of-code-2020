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

func Test_Day08Parse(t *testing.T) {
	testPath := "../../inputs/day08.example.txt"
	expected := Program{
		Ip:  0,
		Acc: 0,
		Commands: []Instruction{
			{Op: "nop", Arg: "+0"},
			{Op: "acc", Arg: "+1"},
			{Op: "jmp", Arg: "+4"},
			{Op: "acc", Arg: "+3"},
			{Op: "jmp", Arg: "-3"},
			{Op: "acc", Arg: "-99"},
			{Op: "acc", Arg: "+1"},
			{Op: "jmp", Arg: "-4"},
			{Op: "acc", Arg: "+6"}},
	}

	data, err := Day08Parse(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_Day08Solution01(t *testing.T) {
	data := Program{
		Ip:  0,
		Acc: 0,
		Commands: []Instruction{
			{Op: "nop", Arg: "+0"},
			{Op: "acc", Arg: "+1"},
			{Op: "jmp", Arg: "+4"},
			{Op: "acc", Arg: "+3"},
			{Op: "jmp", Arg: "-3"},
			{Op: "acc", Arg: "-99"},
			{Op: "acc", Arg: "+1"},
			{Op: "jmp", Arg: "-4"},
			{Op: "acc", Arg: "+6"}},
	}
	expected := "5"
	actual, err := Day08Solution01(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Day08Solution02(t *testing.T) {
	data := Program{
		Ip:  0,
		Acc: 0,
		Commands: []Instruction{
			{Op: "nop", Arg: "+0"},
			{Op: "acc", Arg: "+1"},
			{Op: "jmp", Arg: "+4"},
			{Op: "acc", Arg: "+3"},
			{Op: "jmp", Arg: "-3"},
			{Op: "acc", Arg: "-99"},
			{Op: "acc", Arg: "+1"},
			{Op: "jmp", Arg: "-4"},
			{Op: "acc", Arg: "+6"}},
	}
	expected := "8"
	actual, err := Day08Solution02(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
