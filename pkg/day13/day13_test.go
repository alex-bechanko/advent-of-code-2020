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
package day13

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	data, err := Parse("../../inputs/day13.example.txt")
	expected := solutionData{
		time: big.NewInt(939),
		buses: []*big.Int{
			big.NewInt(7),
			big.NewInt(13),
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(59),
			big.NewInt(0),
			big.NewInt(31),
			big.NewInt(19),
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, data, expected)

}

func Test_Solution1(t *testing.T) {
	data, err := Parse("../../inputs/day13.example.txt")
	assert.NoError(t, err)

	expected := "295"
	actual, err := Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

}

func Test_Solution2(t *testing.T) {
	data, err := Parse("../../inputs/day13.example.txt")
	assert.NoError(t, err)

	expected := "1068781"
	actual, err := Solution2(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
