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
package day11

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const emptySeat = "L"
const occupiedSeat = "#"
const floor = "."

type seatLayout struct {
	width  int
	height int
	layout [][]string
}

func (seats *seatLayout) Copy() seatLayout {
	layout := make([][]string, len(seats.layout))
	for i := range seats.layout {
		layout[i] = make([]string, len(seats.layout[i]))
		copy(layout[i], seats.layout[i])
	}

	return seatLayout{
		width:  seats.width,
		height: seats.height,
		layout: layout,
	}
}

func (seats *seatLayout) String() string {
	sb := strings.Builder{}

	for y := 0; y < seats.height; y++ {
		for x := 0; x < seats.width; x++ {
			sb.WriteString(seats.layout[y][x])
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (seats *seatLayout) equal(other seatLayout) bool {
	if seats.width != other.width || seats.height != other.height {
		return false
	}

	for y := 0; y < seats.height; y++ {
		for x := 0; x < seats.width; x++ {
			if seats.layout[y][x] != other.layout[y][x] {
				return false
			}
		}
	}

	return true
}

func (seats *seatLayout) numOccupiedSeatsAdjacent(x, y int) (int, error) {
	if x >= seats.width || y >= seats.height {
		return 0, fmt.Errorf("Layout coordinates out of bounds: (%d, %d)", x, y)
	}

	coords := [][]int{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
		{x, y + 1},
		{x - 1, y + 1},
		{x - 1, y},
	}
	count := 0
	for _, c := range coords {
		ax := c[0]
		ay := c[1]
		if ax < 0 || ax >= seats.width {
			continue
		}
		if ay < 0 || ay >= seats.height {
			continue
		}

		if seats.layout[ay][ax] == occupiedSeat {
			count++
		}
	}

	return count, nil
}

func (seats *seatLayout) raytraceSeat(x, y, dx, dy int) string {
	//fmt.Printf("Inside raytrace (%d,%d) (%d, %d)\n", x, y, dx, dy)
	for {
		//fmt.Printf("Checking (%d, %d)\n", x+dx, y+dy)
		if !seats.inbounds(x+dx, y+dy) {
			//fmt.Printf("out of bounds, returning floor: (%d, %d)", x+dx, y+dy)
			return floor
		}

		if !seats.floor(x+dx, y+dy) {
			//fmt.Printf("Found seat: (%d, %d)\n", x+dx, y+dy)
			return seats.layout[y+dy][x+dx]
		}

		x += dx
		y += dy
	}
}

func (seats *seatLayout) numOccupiedSeatsRaytrace(x, y int) (int, error) {
	if x >= seats.width || y >= seats.height {
		return 0, fmt.Errorf("Layout coordinates out of bounds: (%d, %d)", x, y)
	}

	raytracedSeats := []string{
		seats.raytraceSeat(x, y, -1, -1), // up-left
		seats.raytraceSeat(x, y, 0, -1),  //up
		seats.raytraceSeat(x, y, +1, -1), //up-right
		seats.raytraceSeat(x, y, +1, 0),  // right
		seats.raytraceSeat(x, y, +1, +1), // down-right
		seats.raytraceSeat(x, y, 0, +1),  // down
		seats.raytraceSeat(x, y, -1, +1), // down-left
		seats.raytraceSeat(x, y, -1, 0),  // left
	}

	count := 0
	for _, tile := range raytracedSeats {
		if tile == occupiedSeat {
			count++
		}
	}

	return count, nil
}

func (seats *seatLayout) inbounds(x, y int) bool {
	if x < 0 || x >= seats.width {
		return false
	} else if y < 0 || y >= seats.height {
		return false
	}

	return true
}

func (seats *seatLayout) updateGrid(x, y int, tile string) {
	if !seats.inbounds(x, y) {
		log.Fatalf("Coords not in bounds: (%d, %d) in (%d, %d) grid", x, y, seats.width, seats.height)
	}

	seats.layout[y][x] = tile
}

func (seats *seatLayout) occupiedSeat(x, y int) bool {
	if !seats.inbounds(x, y) {
		log.Fatalf("Coords not in bounds: (%d, %d) in (%d, %d) grid", x, y, seats.width, seats.height)
	}

	return seats.layout[y][x] == occupiedSeat
}

func (seats *seatLayout) emptySeat(x, y int) bool {
	if !seats.inbounds(x, y) {
		log.Fatalf("Coords not in bounds: (%d, %d) in (%d, %d) grid", x, y, seats.width, seats.height)
	}

	return seats.layout[y][x] == emptySeat
}

func (seats *seatLayout) floor(x, y int) bool {
	if !seats.inbounds(x, y) {
		log.Fatalf("Coords not in bounds: (%d, %d) in (%d, %d) grid", x, y, seats.width, seats.height)
	}

	return seats.layout[y][x] == floor
}

func parseReader(file io.Reader) (seatLayout, error) {
	data := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cs := strings.Split(strings.TrimSpace(line), "")

		data = append(data, cs)
	}

	l := seatLayout{
		layout: data,
		width:  len(data[0]),
		height: len(data),
	}

	return l, nil
}

func Parse(path string) (seatLayout, error) {
	file, err := os.Open(path)
	if err != nil {
		return seatLayout{}, err
	}
	defer file.Close()

	return parseReader(file)

}

func Solution1(seats seatLayout) (string, error) {

	prev := seats.Copy()
	curr := seats.Copy()
	for {
		for x := 0; x < curr.width; x++ {
			for y := 0; y < curr.height; y++ {
				numAdj, err := prev.numOccupiedSeatsAdjacent(x, y)

				if err != nil {
					return "", err
				}

				if numAdj == 0 && prev.emptySeat(x, y) {
					curr.updateGrid(x, y, occupiedSeat)

				} else if numAdj >= 4 && prev.occupiedSeat(x, y) {
					curr.updateGrid(x, y, emptySeat)
				}
			}
		}
		if prev.equal(curr) {
			break
		}
		prev = curr.Copy()
	}

	count := 0
	for x := 0; x < curr.width; x++ {
		for y := 0; y < curr.height; y++ {
			if curr.layout[y][x] == occupiedSeat {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}

func Solution2(seats seatLayout) (string, error) {
	prev := seats.Copy()
	curr := seats.Copy()

	for {
		for x := 0; x < curr.width; x++ {
			for y := 0; y < curr.height; y++ {
				numAdj, err := prev.numOccupiedSeatsRaytrace(x, y)

				if err != nil {
					return "", err
				}

				if numAdj == 0 && prev.emptySeat(x, y) {
					curr.updateGrid(x, y, occupiedSeat)

				} else if numAdj >= 5 && prev.occupiedSeat(x, y) {
					curr.updateGrid(x, y, emptySeat)
				}
			}
		}
		if prev.equal(curr) {
			break
		}
		prev = curr.Copy()
	}

	count := 0
	for x := 0; x < curr.width; x++ {
		for y := 0; y < curr.height; y++ {
			if curr.layout[y][x] == occupiedSeat {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}
