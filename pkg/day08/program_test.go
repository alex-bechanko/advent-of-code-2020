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

func Test_newInstruction(t *testing.T) {
	var inst instruction
	var expected instruction
	var err error

	inst, err = newInstruction("nop +0")
	expected = instruction{op: "nop", arg: "+0"}
	assert.NoError(t, err)
	assert.Equal(t, expected, inst)

	inst, err = newInstruction("acc +1")
	expected = instruction{op: "acc", arg: "+1"}
	assert.NoError(t, err)
	assert.Equal(t, expected, inst)

	inst, err = newInstruction("jmp -3")
	expected = instruction{op: "jmp", arg: "-3"}
	assert.NoError(t, err)
	assert.Equal(t, expected, inst)
}

func Test_programCopy(t *testing.T) {
	prg := program{
		acc: 0,
		ip:  0,
		commands: []instruction{
			{"nop", "+0"},
			{"nop", "+1"},
			{"nop", "-3"},
		},
	}

	expected := program{
		acc: 0,
		ip:  0,
		commands: []instruction{
			{"nop", "+0"},
			{"nop", "+1"},
			{"nop", "-3"},
		},
	}

	assert.Equal(t, prg.Copy(), expected)
}

func Test_programRun(t *testing.T) {
	prg := program{
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

	prg.Run()
	assert.Equal(t, 5, prg.acc)
	assert.Equal(t, 1, prg.ip)
}
