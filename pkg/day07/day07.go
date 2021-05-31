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
package day07

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type BagGraph struct {
	Nodes map[string]bool
	Edges []BagGraphEdge
}

func NewBagGraph() BagGraph {
	return BagGraph{
		Nodes: make(map[string]bool),
		Edges: make([]BagGraphEdge, 0),
	}
}

func (grph *BagGraph) AddNode(color string) {
	grph.Nodes[color] = true
}

func (grph *BagGraph) AddEdge(start, end string, amount int) {
	grph.AddNode(start)
	grph.AddNode(end)
	grph.Edges = append(grph.Edges, BagGraphEdge{
		Start:  start,
		End:    end,
		Weight: amount,
	})
}

type BagGraphEdge struct {
	Start  string
	End    string
	Weight int
}

func ParseFile(path string) (BagGraph, error) {
	bags := NewBagGraph()

	file, err := os.Open(path)
	if err != nil {
		return bags, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.TrimSuffix(line, ".")

		bagDescr := strings.Split(line, "contain")
		start := strings.TrimSpace(bagDescr[0])
		start = strings.TrimSuffix(start, "bags")
		start = strings.TrimSpace(start)
		bags.AddNode(start)

		if strings.TrimSpace(bagDescr[1]) == "no other bags" {
			continue
		}
		edges := strings.Split(strings.TrimSpace(bagDescr[1]), ",")

		for _, e := range edges {
			edge := strings.SplitN(strings.TrimSpace(e), " ", 2)
			end := strings.TrimSpace(edge[1])
			end = strings.TrimSuffix(end, "bags")
			end = strings.TrimSuffix(end, "bag")
			end = strings.TrimSpace(end)

			weight, err := strconv.Atoi(edge[0])
			if err != nil {
				return bags, err
			}
			bags.AddEdge(start, end, weight)
		}
	}

	return bags, nil
}

func Solution1(grph BagGraph) (string, error) {
	targetBag := "shiny gold"

	traveled := make(map[string]bool)

	stack := []string{targetBag}
	for {
		if len(stack) == 0 {
			break
		}

		bag := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, edge := range grph.Edges {
			if _, ok := traveled[edge.Start]; ok {
				// already looked at this bag color, skip
				continue
			} else if edge.End == bag {
				traveled[edge.Start] = true
				stack = append(stack, edge.Start)
			}
		}
	}

	return strconv.Itoa(len(traveled)), nil
}

func Solution2(grph BagGraph) (string, error) {
	stack := []string{"shiny gold"}

	bagCount := -1
	for {
		if len(stack) == 0 {
			break
		}

		bag := stack[len(stack)-1]
		bagCount++
		stack = stack[:len(stack)-1]

		for _, edge := range grph.Edges {
			if edge.Start != bag {
				continue
			}
			for i := 0; i < edge.Weight; i++ {
				stack = append(stack, edge.End)
			}
		}
	}

	return strconv.Itoa(bagCount), nil
}
