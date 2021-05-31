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
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newBagGraph(t *testing.T) {
	bag := newBagGraph()

	assert.Equal(t, 0, len(bag.nodes))
	assert.Equal(t, 0, len(bag.edges))
}

func Test_bagGraphaddNode(t *testing.T) {
	bag := newBagGraph()
	bag.addNode("red")

	assert.Equal(t, 1, len(bag.nodes))
	assert.Equal(t, 0, len(bag.edges))

	yes, ok := bag.nodes["red"]
	assert.True(t, ok)
	assert.True(t, yes)
}

func Test_bagGraphAddEdge(t *testing.T) {
	bag := newBagGraph()
	expectedEdge := bagGraphEdge{start: "red", end: "green", weight: 25}
	bag.addEdge("red", "green", 25)

	assert.Equal(t, 2, len(bag.nodes))
	assert.Equal(t, 1, len(bag.edges))

	var yes bool
	var ok bool

	yes, ok = bag.nodes["red"]
	assert.True(t, ok)
	assert.True(t, yes)

	yes, ok = bag.nodes["green"]
	assert.True(t, ok)
	assert.True(t, yes)

	edge := bag.edges[0]
	assert.Equal(t, expectedEdge, edge)
}
