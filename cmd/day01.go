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
package cmd

import (
	"github.com/alex-bechanko/advent-of-code-2020/pkg/solutions"
	"github.com/spf13/cobra"
)

func init() {
	var fileInput string
	var day01Cmd = &cobra.Command{
		Use:   "day01",
		Short: "Compute the day 1 solutions for Advent of Code 2020",
		Long:  `Computes part1 and part2 solutions for day 1 of Advent of Code 2020`,
		Run: func(cmd *cobra.Command, args []string) {
			solutions.Day01Solutions(&fileInput)
		},
	}

	rootCmd.AddCommand(day01Cmd)
	day01Cmd.Flags().StringVarP(&fileInput, "input", "i", "inputs/day01.txt", "Path to the input data")
}
