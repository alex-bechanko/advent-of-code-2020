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
package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Passenger struct {
	answers map[string]bool
}

func NewPassenger(as string) Passenger {
	p := Passenger{
		answers: make(map[string]bool),
	}

	for _, c := range strings.Split(as, "") {
		p.answers[c] = true
	}

	return p
}

type PlaneGroup struct {
	AnyoneAnswers   map[string]bool
	EveryoneAnswers map[string]bool
	Count           int
}

func NewPlaneGroup() PlaneGroup {
	return PlaneGroup{
		AnyoneAnswers: make(map[string]bool),
		Count:         0,
	}
}

func (grp *PlaneGroup) AddPassenger(p Passenger) {
	for k := range p.answers {
		grp.AnyoneAnswers[k] = true
	}

	grp.Count++
	if grp.Count == 1 {
		grp.EveryoneAnswers = p.answers
	} else {
		answers := make(map[string]bool)
		for k := range grp.EveryoneAnswers {
			if _, ok := p.answers[k]; ok {
				answers[k] = true
			}
		}
		grp.EveryoneAnswers = answers
	}

}

func Day06Parse(path string) ([]PlaneGroup, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	groups := make([]PlaneGroup, 0)
	group := NewPlaneGroup()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			groups = append(groups, group)
			group = NewPlaneGroup()
			continue
		}

		group.AddPassenger(NewPassenger(line))
	}
	groups = append(groups, group)
	return groups, nil
}

func Day06Solution01(plane []PlaneGroup) (string, error) {
	total := 0
	for _, grp := range plane {
		total += len(grp.AnyoneAnswers)
	}
	return strconv.Itoa(total), nil
}

func Day06Solution02(plane []PlaneGroup) (string, error) {
	total := 0
	for _, grp := range plane {
		total += len(grp.EveryoneAnswers)
	}
	return strconv.Itoa(total), nil
}

func Day06Solutions(path *string) {
	plane, err := Day06Parse(*path)
	if err != nil {
		log.Fatal(err)
	}

	soln01, _ := Day06Solution01(plane)
	fmt.Printf("Solution 1: %s\n", soln01)

	soln02, _ := Day06Solution02(plane)
	fmt.Printf("Solution 2: %s\n", soln02)
}
