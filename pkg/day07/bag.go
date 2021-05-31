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

type bagGraph struct {
	nodes map[string]bool
	edges []bagGraphEdge
}

type bagGraphEdge struct {
	start  string
	end    string
	weight int
}

func newBagGraph() bagGraph {
	return bagGraph{
		nodes: make(map[string]bool),
		edges: make([]bagGraphEdge, 0),
	}
}

func (grph *bagGraph) addNode(color string) {
	grph.nodes[color] = true
}

func (grph *bagGraph) addEdge(start, end string, amount int) {
	grph.addNode(start)
	grph.addNode(end)
	grph.edges = append(grph.edges, bagGraphEdge{
		start:  start,
		end:    end,
		weight: amount,
	})
}
