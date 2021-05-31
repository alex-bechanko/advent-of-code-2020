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
package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newSeat(t *testing.T) {
	rs := []intComparison{High, Low, Low, Low, High, High, Low}
	cs := []intComparison{High, High, High}

	expected := seat{
		row:    70,
		column: 7,
		id:     567,
		rows:   []intComparison{High, Low, Low, Low, High, High, Low},
		cols:   []intComparison{High, High, High},
	}

	actual, err := newSeat(rs, cs)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
func Test_findNum(t *testing.T) {
	rs := []intComparison{High, Low, Low, Low, High, High, Low}
	cs := []intComparison{High, High, High}

	var num int
	var err error

	num, err = findNum(0, 127, rs)
	assert.NoError(t, err)
	assert.Equal(t, 70, num)

	num, err = findNum(0, 7, cs)
	assert.NoError(t, err)
	assert.Equal(t, 7, num)
}
