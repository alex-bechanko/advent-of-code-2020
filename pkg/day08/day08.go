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
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ParseFile(path string) (program, error) {
	program := program{
		ip:  0,
		acc: 0,
	}

	file, err := os.Open(path)
	if err != nil {
		return program, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		cmd, err := newInstruction(line)
		if err != nil {
			return program, err
		}
		program.commands = append(program.commands, cmd)

	}

	return program, nil
}

func Solution1(prg program) (string, error) {
	prg.Run()
	return strconv.Itoa(prg.acc), nil
}

func Solution2(start program) (string, error) {
	for i, inst := range start.commands {
		if inst.op == Jmp {
			prg := start.Copy()
			prg.commands[i].op = Nop
			terminated := prg.Run()
			if terminated {
				return strconv.Itoa(prg.acc), nil
			}
		} else if inst.op == Nop {
			prg := start.Copy()
			prg.commands[i].op = Jmp
			terminated := prg.Run()
			if terminated {
				return strconv.Itoa(prg.acc), nil
			}
		}
	}
	return "", fmt.Errorf("no solution")
}
