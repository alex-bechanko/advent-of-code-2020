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
package day12

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	north direction = iota
	south
	east
	west
	forward
	left
	right
)

type move struct {
	direction direction
	magnitude int
}

func (m move) String() string {
	switch m.direction {
	case east:
		return "E" + strconv.Itoa(m.magnitude)
	case west:
		return "W" + strconv.Itoa(m.magnitude)
	case north:
		return "N" + strconv.Itoa(m.magnitude)
	case south:
		return "S" + strconv.Itoa(m.magnitude)
	case left:
		return "L" + strconv.Itoa(m.magnitude)
	case forward:
		return "F" + strconv.Itoa(m.magnitude)
	case right:
		return "R" + strconv.Itoa(m.magnitude)
	default:
		log.Panic("Some how have a move that isn't one of the constants")
	}
	return ""
}

type solutionData struct {
	moves []move
}

func (data solutionData) String() string {
	var b strings.Builder
	for _, m := range data.moves {
		b.WriteString(m.String())
		b.WriteString("\n")
	}

	return b.String()
}

func (s *solutionData) Copy() *solutionData {
	cp := solutionData{
		moves: make([]move, len(s.moves)),
	}

	copy(cp.moves, s.moves)
	return &cp
}

type vec2d struct {
	x float64
	y float64
}

func (v vec2d) rotate(angle float64) vec2d {
	s := math.Sin(angle * math.Pi / 180)
	c := math.Cos(angle * math.Pi / 180)

	nv := vec2d{}
	nv.x = c*v.x - s*v.y
	nv.y = s*v.x + c*v.y
	return nv
}

func (v vec2d) add(x, y float64) vec2d {
	return vec2d{
		x: v.x + x,
		y: v.y + y,
	}
}

func (v vec2d) manhatten() float64 {
	x := v.x
	if x < 0 {
		x *= -1
	}
	y := v.y
	if y < 0 {
		y *= -1
	}

	return x + y
}

func Parse(path string) (solutionData, error) {
	file, err := os.Open(path)
	if err != nil {
		return solutionData{}, err
	}
	defer file.Close()

	mapping := map[string]direction{
		"F": forward,
		"L": left,
		"R": right,
		"N": north,
		"S": south,
		"E": east,
		"W": west,
	}

	data := make([]move, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cs := strings.SplitN(strings.TrimSpace(line), "", 2)

		dir, ok := mapping[cs[0]]
		if !ok {
			return solutionData{}, fmt.Errorf("Invalid direction found while parsing: %s", cs[0])
		}

		mag, err := strconv.Atoi(cs[1])
		if err != nil {
			return solutionData{}, fmt.Errorf("Invalid magnitude found while parsing: %s", cs[1])
		}

		data = append(data, move{direction: dir, magnitude: mag})
	}

	return solutionData{moves: data}, nil
}

func Solution1(data solutionData) (string, error) {
	x := 0
	y := 0
	angle := 0

	for _, m := range data.moves {
		switch m.direction {
		case north:
			y += m.magnitude
		case south:
			y -= m.magnitude
		case east:
			x += m.magnitude
		case west:
			x -= m.magnitude
		case left:
			angle += m.magnitude
		case right:
			angle -= m.magnitude
		case forward:
			x += m.magnitude * int(math.Cos(float64(angle)/180*math.Pi))
			y += m.magnitude * int(math.Sin(float64(angle)/180*math.Pi))
		default:
			return "", fmt.Errorf("Invalid direction provided")
		}
	}

	return strconv.Itoa(int(math.Abs(float64(x))) + int(math.Abs(float64(y)))), nil
}

func Solution2(data solutionData) (string, error) {
	boat := vec2d{x: 0, y: 0}
	vel := vec2d{x: 10, y: 1}

	for _, m := range data.moves {
		mag := float64(m.magnitude)
		switch m.direction {
		case north:
			vel = vel.add(0, mag)
		case south:
			vel = vel.add(0, -1*mag)
		case east:
			vel = vel.add(mag, 0)
		case west:
			vel = vel.add(-1*mag, 0)
		case left:
			vel = vel.rotate(mag)
		case right:
			vel = vel.rotate(-1 * mag)
		case forward:
			boat = boat.add(vel.x*mag, vel.y*mag)
		default:
			return "", fmt.Errorf("Invalid direction provided")
		}
	}

	return strconv.Itoa(int(boat.manhatten())), nil
}
