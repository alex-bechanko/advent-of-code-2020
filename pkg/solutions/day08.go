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
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type InstrType string

const (
	Nop = "nop"
	Acc = "acc"
	Jmp = "jmp"
)

type Instruction struct {
	Op  string
	Arg string
}

type Program struct {
	Ip       int
	Acc      int
	Commands []Instruction
}

func NewInstruction(line string) (Instruction, error) {
	inst := Instruction{}

	args := strings.SplitN(line, " ", 2)
	if len(args) != 2 {
		return inst, fmt.Errorf("Instruction string could not be broken into two pieces: \"%s\"", line)
	}

	inst.Op = strings.TrimSpace(args[0])
	inst.Arg = strings.TrimSpace(args[1])

	return inst, nil
}

func Day08Parse(path string) (Program, error) {
	program := Program{
		Ip:  0,
		Acc: 0,
	}

	file, err := os.Open(path)
	if err != nil {
		return program, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		cmd, err := NewInstruction(line)
		if err != nil {
			return program, err
		}
		program.Commands = append(program.Commands, cmd)

	}

	return program, nil
}

func (prg *Program) Copy() Program {
	p := Program{
		Ip:  prg.Ip,
		Acc: prg.Acc,
	}
	p.Commands = make([]Instruction, len(prg.Commands))
	copy(p.Commands, prg.Commands)
	return p
}

func (prg *Program) Run() bool {
	runCmds := make(map[int]bool)

	for {
		if _, ok := runCmds[prg.Ip]; ok {
			return false
		}
		if len(prg.Commands) <= prg.Ip {
			return true
		}

		cmd := prg.Commands[prg.Ip]
		runCmds[prg.Ip] = true
		switch cmd.Op {
		case Nop:
			prg.Ip++
		case Acc:
			prg.Ip++
			acc, err := strconv.Atoi(cmd.Arg)
			if err != nil {
				log.Fatal(err)
			}
			prg.Acc += acc
		case Jmp:
			inc, err := strconv.Atoi(cmd.Arg)
			if err != nil {
				log.Fatal(err)
			}
			prg.Ip += inc
		}
	}
}

func Day08Solution01(prg Program) (string, error) {
	prg.Run()
	return strconv.Itoa(prg.Acc), nil
}

func Day08Solution02(start Program) (string, error) {
	for i, inst := range start.Commands {
		if inst.Op == Jmp {
			prg := start.Copy()
			prg.Commands[i].Op = Nop
			terminated := prg.Run()
			if terminated {
				return strconv.Itoa(prg.Acc), nil
			}
		} else if inst.Op == Nop {
			prg := start.Copy()
			prg.Commands[i].Op = Jmp
			terminated := prg.Run()
			if terminated {
				return strconv.Itoa(prg.Acc), nil
			}
		}
	}
	return "", fmt.Errorf("no solution")
}

func Day08Solutions(path *string) {
	prg, err := Day08Parse(*path)
	if err != nil {
		log.Fatal(err)
	}

	soln01, _ := Day08Solution01(prg)
	fmt.Printf("Solution 1: %s\n", soln01)

	soln02, _ := Day08Solution02(prg)
	fmt.Printf("Solution 2: %s\n", soln02)
}
