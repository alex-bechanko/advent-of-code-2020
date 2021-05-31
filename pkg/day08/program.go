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
	"fmt"
	"log"
	"strconv"
	"strings"
)

type InstrType string

const (
	Nop = "nop"
	Acc = "acc"
	Jmp = "jmp"
)

type instruction struct {
	op  string
	arg string
}

type program struct {
	ip       int
	acc      int
	commands []instruction
}

func newInstruction(line string) (instruction, error) {
	inst := instruction{}

	args := strings.SplitN(line, " ", 2)
	if len(args) != 2 {
		return inst, fmt.Errorf("Instruction string could not be broken into two pieces: \"%s\"", line)
	}

	inst.op = strings.TrimSpace(args[0])
	inst.arg = strings.TrimSpace(args[1])

	return inst, nil
}

func (prg *program) Copy() program {
	p := program{
		ip:  prg.ip,
		acc: prg.acc,
	}
	p.commands = make([]instruction, len(prg.commands))
	copy(p.commands, prg.commands)
	return p
}

func (prg *program) Run() bool {
	runCmds := make(map[int]bool)

	for {
		if _, ok := runCmds[prg.ip]; ok {
			return false
		}
		if len(prg.commands) <= prg.ip {
			return true
		}

		cmd := prg.commands[prg.ip]
		runCmds[prg.ip] = true
		switch cmd.op {
		case Nop:
			prg.ip++
		case Acc:
			prg.ip++
			acc, err := strconv.Atoi(cmd.arg)
			if err != nil {
				log.Fatal(err)
			}
			prg.acc += acc
		case Jmp:
			inc, err := strconv.Atoi(cmd.arg)
			if err != nil {
				log.Fatal(err)
			}
			prg.ip += inc
		}
	}
}
