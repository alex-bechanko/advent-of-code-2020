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
	"strings"
)

type passenger struct {
	answers map[string]bool
}

type planeGroup struct {
	anyoneAnswers   map[string]bool
	everyoneAnswers map[string]bool
	count           int
}

func newPassenger(as string) passenger {
	p := passenger{
		answers: make(map[string]bool),
	}

	for _, c := range strings.Split(as, "") {
		p.answers[c] = true
	}

	return p
}

func newPlaneGroup() planeGroup {
	return planeGroup{
		anyoneAnswers: make(map[string]bool),
		count:         0,
	}
}

func (grp *planeGroup) addPassenger(p passenger) {
	for k := range p.answers {
		grp.anyoneAnswers[k] = true
	}

	grp.count++
	if grp.count == 1 {
		grp.everyoneAnswers = p.answers
	} else {
		answers := make(map[string]bool)
		for k := range grp.everyoneAnswers {
			if _, ok := p.answers[k]; ok {
				answers[k] = true
			}
		}
		grp.everyoneAnswers = answers
	}

}
