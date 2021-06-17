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
package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intstack []int

type program struct {
	mask         *mask
	memory       map[int]int
	instructions []instruction
}

type instruction interface {
	runV1(*program)
	runV2(*program)
}

type mask struct {
	floats []int
	ones   []int
	zeros  []int
}

type memset struct {
	addr  int
	value int
}

func parseMask(m string) (mask, error) {
	floats := make([]int, 0)
	ones := make([]int, 0)
	zeros := make([]int, 0)

	m = strings.TrimSpace(m)
	if !strings.HasPrefix(m, "mask = ") {
		return mask{}, fmt.Errorf("unable to parse mask %s", m)
	}

	ms := strings.Split(strings.TrimPrefix(m, "mask = "), "")
	for i, c := range ms {
		switch c {
		case "1":
			ones = append(ones, len(ms)-1-i)
		case "0":
			zeros = append(zeros, len(ms)-1-i)
		case "X":
			floats = append(floats, len(ms)-1-i)
		default:
			return mask{}, fmt.Errorf("unable to parse mask %s", m)
		}
	}

	return mask{ones: ones, zeros: zeros, floats: floats}, nil
}

func parseMemset(m string) (memset, error) {
	m = strings.TrimSpace(m)
	line := m
	if !strings.HasPrefix(m, "mem[") {
		return memset{}, fmt.Errorf("unable to parse memset %s", line)
	}
	m = strings.TrimPrefix(m, "mem[")

	ms := strings.Split(m, "]")
	if len(ms) != 2 {
		return memset{}, fmt.Errorf("unable to parse memset %s", line)
	}

	n, m := ms[0], strings.TrimPrefix(ms[1], " = ")
	addr, err := strconv.Atoi(n)
	if err != nil {
		return memset{}, fmt.Errorf("unable to parse memset %s: %s", line, err)
	}

	value, err := strconv.Atoi(m)
	if err != nil {
		return memset{}, fmt.Errorf("unable to parse memset %s: %s", line, err)
	}

	return memset{addr: addr, value: value}, nil

}

func (p *program) runV1() {
	for _, inst := range p.instructions {
		inst.runV1(p)
	}
}

func (m *mask) runV1(p *program) {
	p.mask = m
}

func (m *memset) runV1(p *program) {
	num := m.value
	for _, o := range p.mask.ones {
		num |= (1 << o)
	}

	for _, z := range p.mask.zeros {
		num &= ^(1 << z)
	}

	p.memory[m.addr] = num
}

func (p *program) runV2() {
	for _, inst := range p.instructions {
		inst.runV2(p)
	}
}

func (m *mask) runV2(p *program) {
	p.mask = m
}

func (m *memset) runV2(p *program) {
	newAddr := m.addr
	for _, o := range p.mask.ones {
		newAddr |= (1 << o)
	}

	for i := 0; i < (1 << len(p.mask.floats)); i++ {
		mem := newAddr
		// each int i represents the set of floats

		for j, f := range p.mask.floats {
			if (i>>j)&1 == 1 {
				// the jth bit in i is set to 1, so flip the memory bit to one
				mem |= (1 << f)
			} else {
				// jth bit in i is 0, so memory bit should be zero
				mem &= ^(1 << f)
			}
		}

		// after inner loop completes we have a single combination of floats that need to be set
		p.memory[mem] = m.value

	}

}

func Parse(path string) (program, error) {
	file, err := os.Open(path)
	if err != nil {
		return program{}, err
	}
	defer file.Close()

	instr := make([]instruction, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "mask") {
			m, err := parseMask(line)
			if err != nil {
				return program{}, err
			}
			instr = append(instr, &m)
		} else if strings.HasPrefix(line, "mem") {
			m, err := parseMemset(line)
			if err != nil {
				return program{}, err
			}
			instr = append(instr, &m)
		}
	}
	return program{
		mask:         nil,
		memory:       make(map[int]int),
		instructions: instr,
	}, nil
}

func Solution1(program program) (string, error) {
	program.runV1()
	total := 0
	for _, v := range program.memory {
		total += v
	}

	return strconv.Itoa(total), nil
}
func Solution2(program program) (string, error) {
	program.runV2()
	total := 0
	for _, v := range program.memory {
		total += v
	}

	return strconv.Itoa(total), nil
}
