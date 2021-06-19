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
package day17

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	x int
	y int
	z int
}

func (v Vertex) Neighbors() []Vertex {
	// in a cube environment, number of neighbors is always 26
	const numNeighbors int = 26

	ns := make([]Vertex, numNeighbors)
	ind := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				ns[ind] = Vertex{
					x: v.x + i,
					y: v.y + j,
					z: v.z + k,
				}
				ind++

			}
		}
	}

	return ns
}

func (v Vertex) NumActiveNeighbors(vs map[Vertex]bool) int {
	count := 0
	for _, n := range v.Neighbors() {
		if active, ok := vs[n]; ok && active {
			count++
		}
	}
	return count
}

type HyperVertex struct {
	x int
	y int
	z int
	w int
}

func (v HyperVertex) Neighbors() []HyperVertex {
	// in a cube environment, number of neighbors is always 26
	const numNeighbors int = 80

	ns := make([]HyperVertex, numNeighbors)
	ind := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for m := -1; m <= 1; m++ {

					if i == 0 && j == 0 && k == 0 && m == 0 {
						continue
					}
					ns[ind] = HyperVertex{
						x: v.x + i,
						y: v.y + j,
						z: v.z + k,
						w: v.w + m,
					}
					ind++
				}
			}
		}
	}

	return ns
}

func (v HyperVertex) NumActiveNeighbors(vs map[HyperVertex]bool) int {
	count := 0
	for _, n := range v.Neighbors() {
		if active, ok := vs[n]; ok && active {
			count++
		}
	}
	return count
}

type PuzzleInput struct {
	Cubes map[Vertex]bool
}

func (puzzle PuzzleInput) NextGeneration() PuzzleInput {
	nextGen := make(map[Vertex]bool)
	// loop through all vertices defined in Cubes and all the neighbors of the vertices in Cubes
	for v, active := range puzzle.AllAdjacentVertices() {
		num := v.NumActiveNeighbors(puzzle.Cubes)
		if active && (num == 2 || num == 3) {
			nextGen[v] = true
		} else if !active && num == 3 {
			nextGen[v] = true
		}
	}

	return PuzzleInput{Cubes: nextGen}
}

// AllAdjacentVertices returns a map of the cubes adjacent to the active cubes
// including the active cubes. The map stores false for inactive cubes,
// and true for active cubes
func (puzzle PuzzleInput) AllAdjacentVertices() map[Vertex]bool {
	adj := make(map[Vertex]bool)

	for v := range puzzle.Cubes {
		adj[v] = true
		for _, n := range v.Neighbors() {
			_, ok := puzzle.Cubes[n]
			adj[n] = ok
		}
	}

	return adj
}

func (puzzle PuzzleInput) ToHyperPuzzleInput() HyperPuzzleInput {
	cubes := make(map[HyperVertex]bool)
	for v, _ := range puzzle.Cubes {
		h := HyperVertex{x: v.x, y: v.y, z: v.z, w: 0}
		cubes[h] = true
	}

	return HyperPuzzleInput{Cubes: cubes}
}

type HyperPuzzleInput struct {
	Cubes map[HyperVertex]bool
}

// AllAdjacentVertices returns a map of the cubes adjacent to the active cubes
// including the active cubes. The map stores false for inactive cubes,
// and true for active cubes
func (puzzle HyperPuzzleInput) AllAdjacentVertices() map[HyperVertex]bool {
	adj := make(map[HyperVertex]bool)

	for v := range puzzle.Cubes {
		adj[v] = true
		for _, n := range v.Neighbors() {
			_, ok := puzzle.Cubes[n]
			adj[n] = ok
		}
	}

	return adj
}

func (puzzle HyperPuzzleInput) NextGeneration() HyperPuzzleInput {
	nextGen := make(map[HyperVertex]bool)
	// loop through all vertices defined in Cubes and all the neighbors of the vertices in Cubes
	for v, active := range puzzle.AllAdjacentVertices() {
		num := v.NumActiveNeighbors(puzzle.Cubes)
		if active && (num == 2 || num == 3) {
			nextGen[v] = true
		} else if !active && num == 3 {
			nextGen[v] = true
		}
	}

	return HyperPuzzleInput{Cubes: nextGen}
}

func Parse(path string) (PuzzleInput, error) {

	data := make(map[Vertex]bool)
	// all vertexes in the input have z=0
	z := 0

	file, err := os.Open(path)
	if err != nil {
		return PuzzleInput{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		cubes := strings.Split(line, "")
		for x, c := range cubes {
			if c == "#" {
				data[Vertex{x: x, y: y, z: z}] = true
			}
		}
		y++
	}

	return PuzzleInput{Cubes: data}, nil
}

func Solution1(puzzle PuzzleInput) (string, error) {
	for i := 0; i < 6; i++ {
		puzzle = puzzle.NextGeneration()
	}

	return strconv.Itoa(len(puzzle.Cubes)), nil
}

func Solution2(puzzle PuzzleInput) (string, error) {

	hyperPuzzle := puzzle.ToHyperPuzzleInput()
	for i := 0; i < 6; i++ {
		hyperPuzzle = hyperPuzzle.NextGeneration()
	}

	return strconv.Itoa(len(hyperPuzzle.Cubes)), nil
}
