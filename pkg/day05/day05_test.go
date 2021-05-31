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

func Test_ParseFile(t *testing.T) {
	testPath := "../../inputs/day05.example.txt"
	expected := []seat{
		{
			row:    70,
			column: 7,
			id:     567,
			rows:   []intComparison{High, Low, Low, Low, High, High, Low},
			cols:   []intComparison{High, High, High},
		},
		{
			row:    14,
			column: 7,
			id:     119,
			rows:   []intComparison{Low, Low, Low, High, High, High, Low},
			cols:   []intComparison{High, High, High},
		},
		{
			row:    102,
			column: 4,
			id:     820,
			rows:   []intComparison{High, High, Low, Low, High, High, Low},
			cols:   []intComparison{High, Low, Low},
		},
	}
	data, err := ParseFile(testPath)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func Test_Solution1(t *testing.T) {
	data := []seat{
		{
			row:    70,
			column: 7,
			id:     567,
			rows:   []intComparison{High, Low, Low, Low, High, High, Low},
			cols:   []intComparison{High, High, High},
		},
		{
			row:    14,
			column: 7,
			id:     119,
			rows:   []intComparison{Low, Low, Low, High, High, High, Low},
			cols:   []intComparison{High, High, High},
		},
		{
			row:    102,
			column: 4,
			id:     820,
			rows:   []intComparison{High, High, Low, Low, High, High, Low},
			cols:   []intComparison{High, Low, Low},
		},
	}
	expected := "820"
	actual, err := Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Solution2(t *testing.T) {
	data := []seat{
		{
			row:    70,
			column: 7,
			id:     567,
			rows:   []intComparison{High, Low, Low, Low, High, High, Low},
			cols:   []intComparison{High, High, High},
		},
		{
			row:    14,
			column: 7,
			id:     119,
			rows:   []intComparison{Low, Low, Low, High, High, High, Low},
			cols:   []intComparison{High, High, High},
		},
		{
			row:    102,
			column: 4,
			id:     820,
			rows:   []intComparison{High, High, Low, Low, High, High, Low},
			cols:   []intComparison{High, Low, Low},
		},
	}

	_, err := Solution2(data)
	assert.Error(t, err)
}
