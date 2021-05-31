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
package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution1(t *testing.T) {
	data := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	actual, err := Solution1(data)
	assert.NoError(t, err)
	assert.Equal(t, "35", actual)
}

func Test_Solution2(t *testing.T) {
	data := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	actual, err := Solution2(data)
	assert.NoError(t, err)
	assert.Equal(t, "8", actual)
}
