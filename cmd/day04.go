/*
Copyright © 2021 Alex Bechanko

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
	"fmt"
	"os"

	"github.com/alex-bechanko/advent-of-code-2020/pkg/day04"
	"github.com/spf13/cobra"
)

func init() {

	var defaultInput = "inputs/day04.txt"
	var fileInput string

	var day04Cmd = &cobra.Command{
		Use:   "day04",
		Short: "Compute day 4 solutions",
		Long:  "Computes part1 and part2 solutions for day 4 of Advent of Code 2020",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := day04.ParseFile(fileInput)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing data file: %v\n", err)
				os.Exit(1)
			}

			soln1, err := day04.Solution1(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error computing solution 1: %v", err)
			}

			soln2, err := day04.Solution2(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error computing solution 2: %v", err)
			}

			fmt.Printf("Solution1: %s\n", soln1)
			fmt.Printf("Solution2: %s\n", soln2)
		},
	}

	day04Cmd.Flags().StringVarP(&fileInput, "input", "i", defaultInput, "Path to the input data")

	rootCmd.AddCommand(day04Cmd)
}
