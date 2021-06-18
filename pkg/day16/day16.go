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
package day16

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type RuleRange struct {
	Low  int
	High int
}

type Rule struct {
	Name   string
	Range1 RuleRange
	Range2 RuleRange
}

type PuzzleInput struct {
	MyTicket      []int
	NearbyTickets [][]int
	Rules         map[string]Rule
}

func ParseCommaDelimitedInts(data string) ([]int, error) {
	ns := strings.Split(data, ",")

	nums := make([]int, len(ns))
	for i, s := range ns {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, fmt.Errorf("error parsing number in comma delimited list of ints '%s': %s", s, err)
		}

		nums[i] = n
	}

	return nums, nil
}

func parseRange(data string) (RuleRange, error) {
	r := RuleRange{}

	strs := strings.Split(data, "-")
	if len(strs) != 2 {
		return r, fmt.Errorf("invalid number of values parsed from rule range: %s", data)
	}

	low, err := strconv.Atoi(strs[0])
	if err != nil {
		return r, fmt.Errorf("failed to parse rule range's number '%s': %s", strs[0], err)
	}

	high, err := strconv.Atoi(strs[1])
	if err != nil {
		return r, fmt.Errorf("failed to parse rule range's number '%s': %s", strs[1], err)
	}

	r.Low = low
	r.High = high
	return r, nil
}

func parseRule(line string) (Rule, error) {
	buf := bufio.NewReader(strings.NewReader(line))

	data := Rule{}

	name, err := buf.ReadString(':')
	name = strings.TrimSuffix(name, ":")
	if err != nil {
		return data, fmt.Errorf("error parsing rule name from line: %s", err)
	}

	// discard the first space after the colon (:)
	if _, err = buf.Discard(1); err != nil {
		return data, fmt.Errorf("error discarding ':' from line: %s", err)
	}

	r1, err := buf.ReadString(' ')
	r1 = strings.TrimSuffix(r1, " ")
	if err != nil {
		return data, fmt.Errorf("error parsing first rule range from line: %s", err)
	}

	range1, err := parseRange(r1)
	if err != nil {
		return data, fmt.Errorf("error parsing first rule range: %s", err)
	}

	// discard the string "or " to get to the next part of the range
	if _, err = buf.Discard(3); err != nil {
		return data, fmt.Errorf("error discarding 'or ' from line: %s", err)
	}

	r2bytes, err := ioutil.ReadAll(buf)
	r2 := strings.TrimSuffix(string(r2bytes), "\n")
	if err != nil {
		return data, fmt.Errorf("error parsing second rule range from line: %s", err)
	}

	range2, err := parseRange(r2)
	if err != nil {
		return data, fmt.Errorf("error parsing second rule range: %s", err)
	}

	data.Name = name
	data.Range1 = range1
	data.Range2 = range2

	return data, nil
}

func Parse(path string) (PuzzleInput, error) {

	data := PuzzleInput{}

	file, err := os.Open(path)
	if err != nil {
		return data, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string

	rules := make(map[string]Rule)
	// start with parsing rules
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			break
		}

		rule, err := parseRule(line)
		if err != nil {
			return data, fmt.Errorf("error parsing rule from line '%s': %s", line, err)
		}
		rules[rule.Name] = rule
	}

	scanner.Scan()
	line = scanner.Text()
	if line != "your ticket:" {
		return data, fmt.Errorf("error parsing file, unexpected line '%s' when expecting 'your ticket:'", line)
	}

	scanner.Scan()
	line = scanner.Text()
	myTicket, err := ParseCommaDelimitedInts(line)
	if err != nil {
		return data, fmt.Errorf("error parsing my ticket '%s': %s", line, err)
	}

	scanner.Scan()
	line = scanner.Text()
	if line != "" {
		return data, fmt.Errorf("error parsing file, unexpected line '%s' where empty line expected", line)
	}

	scanner.Scan()
	line = scanner.Text()
	if line != "nearby tickets:" {
		return data, fmt.Errorf("error parsing file, unexpected line '%s' when expecting 'nearby tickets:'", line)
	}

	//rest of the lines are for nearby tickets
	nearbyTickets := make([][]int, 0)
	for scanner.Scan() {
		line = scanner.Text()

		ticket, err := ParseCommaDelimitedInts(line)
		if err != nil {
			return data, fmt.Errorf("error parsing nearby ticket '%s': %s", line, err)
		}
		nearbyTickets = append(nearbyTickets, ticket)
	}

	data.Rules = rules
	data.MyTicket = myTicket
	data.NearbyTickets = nearbyTickets

	return data, nil

}

func valuePassesRule(value int, rule Rule) bool {
	if value >= rule.Range1.Low && value <= rule.Range1.High {
		return true
	} else if value >= rule.Range2.Low && value <= rule.Range2.High {
		return true
	}

	return false
}

func invalidValues(ticket []int, rules map[string]Rule) []int {

	invalidValues := make([]int, 0)
	for _, val := range ticket {
		validValue := false

		for _, rule := range rules {
			if valuePassesRule(val, rule) {
				validValue = true
				break
			}
		}

		if !validValue {
			invalidValues = append(invalidValues, val)
		}
	}

	return invalidValues
}

func Solution1(data PuzzleInput) (string, error) {

	total := 0
	// loop through nearby tickets
	for _, ticket := range data.NearbyTickets {
		invalids := invalidValues(ticket, data.Rules)
		for _, v := range invalids {
			total += v
		}
	}

	return strconv.Itoa(total), nil
}

func Solution2(data PuzzleInput) (string, error) {

	possibleRules := make([]map[string]bool, len(data.NearbyTickets[0]))
	for i := range possibleRules {
		rules := make(map[string]bool)
		for _, r := range data.Rules {
			rules[r.Name] = true
		}
		possibleRules[i] = rules
	}

	for _, ticket := range data.NearbyTickets {
		invs := invalidValues(ticket, data.Rules)
		if len(invs) > 0 {
			continue
		}

		// narrow down possible fields by filtering the matching rules to the ticket list
		for i, val := range ticket {
			validRules := make(map[string]bool)

			for r := range possibleRules[i] {
				if valuePassesRule(val, data.Rules[r]) {
					validRules[r] = true
				}
			}

			possibleRules[i] = validRules
		}
	}

	// if a position only has one possible field, then all other positions cannot be this field.
	// so remove that field from the other positions if it is there.
	// repeat until no changes to possible positions

	for change := true; change; {
		change = false

		for i, r1 := range possibleRules {
			if len(r1) > 1 {
				continue
			}

			rule := ""
			for r := range r1 {
				rule = r
			}

			for j, r2 := range possibleRules {
				if i == j {
					continue
				}

				if _, ok := r2[rule]; ok {
					change = true
					delete(r2, rule)
				}

			}

		}
	}

	// at this point the fields should all be narrowed down to one field per position
	// if it is not then that implies the data is bad by inputing cycles in the data
	// validate to ensure all positions have been narrowed down
	fields := make([]string, 0)
	for i, v := range possibleRules {
		if len(v) != 1 {
			return "", fmt.Errorf("a cycle must have occurred for the %d position", i)
		}

		// I hate that this is how I must extract a single unknown key from a map
		field := ""
		for k := range v {
			field = k
		}

		fields = append(fields, field)
	}

	answer := 1
	for i, field := range fields {
		if strings.HasPrefix(field, "departure") {
			answer *= data.MyTicket[i]
		}
	}

	return strconv.Itoa(answer), nil
}
