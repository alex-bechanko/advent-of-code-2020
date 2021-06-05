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
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type solutionData struct {
	time  *big.Int
	buses []*big.Int
}

func (soln solutionData) Copy() solutionData {
	cp := solutionData{
		time:  soln.time,
		buses: make([]*big.Int, len(soln.buses)),
	}
	copy(cp.buses, soln.buses)
	return cp
}

func Parse(path string) (solutionData, error) {
	file, err := os.Open(path)
	if err != nil {
		return solutionData{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buses := make([]*big.Int, 0)
	time := big.NewInt(0)
	scanner.Scan()
	timeLine := scanner.Text()
	t, err := strconv.ParseInt(strings.TrimSpace(timeLine), 10, 64)
	if err != nil {
		return solutionData{}, err
	}
	time.SetInt64(t)

	for scanner.Scan() {
		line := scanner.Text()

		cs := strings.Split(strings.ReplaceAll(strings.TrimSpace(line), "x", "0"), ",")
		for _, c := range cs {
			bus, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				return solutionData{}, err
			}

			buses = append(buses, big.NewInt(bus))
		}
	}

	return solutionData{time: time, buses: buses}, nil
}

func Solution1(data solutionData) (string, error) {
	zero := big.NewInt(0)
	one := big.NewInt(1)

	var minTime, ans *big.Int
	minTime = big.NewInt(0)
	ans = big.NewInt(0)

	var r, q big.Int // used to store divmod in loop
	for i, b := range data.buses {
		if b.Cmp(zero) == 0 {
			continue
		}

		q.DivMod(data.time, b, &r)
		if r.Cmp(zero) == 1 {
			q.Add(&q, one)
		}
		q.Mul(&q, b)

		if minTime.Cmp(&q) == 1 || minTime.Cmp(zero) == 0 {
			minTime.Set(&q)
			ans.Set(ans.Mul(ans.Sub(minTime, data.time), data.buses[i]))
		}
	}

	return ans.String(), nil
}

func Solution2(data solutionData) (string, error) {
	zero := big.NewInt(0)
	as := make([]*big.Int, 0)
	ns := make([]*big.Int, 0)
	for i, b := range data.buses {
		if b.Cmp(zero) == 0 {
			continue
		}
		var a big.Int
		a.Sub(b, big.NewInt(int64(i)))
		as = append(as, &a)
		ns = append(ns, b)
	}

	// confirm that numbers are coprime with each other
	// while we are at it, compute the product of the numbers
	one := big.NewInt(1)
	N := big.NewInt(1)

	for i := 0; i < len(ns); i++ {
		N.Mul(N, ns[i])
	}

	var z big.Int  // used to store gcd(a,b)
	var Ni big.Int // used to store N / ns[i]
	var y big.Int  // used to store the number y where ax + by = gcd(a,b)
	var ans big.Int
	for i, ni := range ns {
		fmt.Println(N, ni)
		Ni.Div(N, ni)
		z.GCD(nil, &y, ni, &Ni)
		if z.Cmp(one) != 0 {
			return "", fmt.Errorf("Can't compute solution, bus intervals are not coprime")
		}

		ans.Add(&ans, y.Mul(as[i], y.Mul(&y, &Ni)))
	}
	ans.Mod(&ans, N)

	return ans.String(), nil
}
