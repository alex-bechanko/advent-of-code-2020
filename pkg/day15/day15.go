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
package day15

import (
	"strconv"

	"github.com/alex-bechanko/advent-of-code-2020/pkg/common"
)

type MemoryGame struct {
	// map from number spoken to when it was last spoken
	Spoken          map[int]int
	StartingNumbers []int

	// retain what the last spoken was separately too for easier calculation
	LastSpoken int
}

func Parse(path string) (MemoryGame, error) {
	nums, err := common.ParseIntFile(path)
	if err != nil {
		return MemoryGame{}, err
	}

	spoken := make(map[int]int)

	g := MemoryGame{
		Spoken:          spoken,
		LastSpoken:      0,
		StartingNumbers: nums,
	}

	return g, nil

}

func PlayGame(data MemoryGame, endTurn int) (string, error) {
	if endTurn < len(data.StartingNumbers) {
		return strconv.Itoa(data.StartingNumbers[endTurn-1]), nil
	}

	for i, n := range data.StartingNumbers {
		if i < len(data.StartingNumbers)-1 {
			data.Spoken[n] = i + 1
		}
	}
	// last element in starting numbers is the first last spoken value, keep it out of map
	// for the algorithm below
	data.LastSpoken = data.StartingNumbers[len(data.StartingNumbers)-1]

	for i := len(data.StartingNumbers); i < endTurn; i++ {
		v, ok := data.Spoken[data.LastSpoken]

		if !ok {
			// number is not in map, so it is a new number
			data.Spoken[data.LastSpoken] = i
			data.LastSpoken = 0
		} else {
			// not new number, so update the turn for the old number
			// and last spoken as the turn difference
			data.Spoken[data.LastSpoken] = i
			data.LastSpoken = i - v
		}
	}

	return strconv.Itoa(data.LastSpoken), nil
}

func Solution1(data MemoryGame) (string, error) {
	return PlayGame(data, 2020)
}

func Solution2(data MemoryGame) (string, error) {
	return PlayGame(data, 30000000)
}
