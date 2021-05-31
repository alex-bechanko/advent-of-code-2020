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
package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newPassenger(t *testing.T) {
	data := "abc"
	expected := passenger{
		answers: map[string]bool{"a": true, "b": true, "c": true},
	}

	assert.Equal(t, expected, newPassenger(data))

}

func Test_groupAddPassenger(t *testing.T) {
	grp := newPlaneGroup()
	grp.addPassenger(newPassenger("ab"))
	grp.addPassenger(newPassenger("ac"))

	expected := planeGroup{
		everyoneAnswers: map[string]bool{"a": true},
		anyoneAnswers:   map[string]bool{"a": true, "b": true, "c": true},
		count:           2,
	}

	assert.Equal(t, expected, grp)
}
