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
package main

import (
	"fmt"
	"os"

	"github.com/alex-bechanko/advent-of-code-2020/pkg/solutions"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func makeDayCommand(root *cobra.Command, day int, daySolutionFunc func(*string)) {
	var fileInput string

	cmdText := fmt.Sprintf("day%02d", day)
	shortText := fmt.Sprintf("Compute the day %d solutions for Advent of Code 2020", day)
	longText := fmt.Sprintf(`Computes part1 and part2 solutions for day %d of Advent of Code 2020`, day)

	var dayCmd = &cobra.Command{
		Use:   cmdText,
		Short: shortText,
		Long:  longText,
		Run: func(cmd *cobra.Command, args []string) {
			daySolutionFunc(&fileInput)
		},
	}
	root.AddCommand(dayCmd)

	var defaultInput = fmt.Sprintf("inputs/day%02d.txt", day)
	dayCmd.Flags().StringVarP(&fileInput, "input", "i", defaultInput, "Path to the input data")
}

var cfgFile string

func main() {

	var rootCmd = &cobra.Command{
		Use:   "advent-of-code-2020",
		Short: "Compute solutions for Advent of Code 2020",
		Long:  `Compute solutions for days 1-25 of Advent of Code 2020`,
	}

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent-of-code-2020.yaml)")

	makeDayCommand(rootCmd, 1, solutions.Day01Solutions)
	makeDayCommand(rootCmd, 2, solutions.Day02Solutions)
	makeDayCommand(rootCmd, 3, solutions.Day03Solutions)
	makeDayCommand(rootCmd, 4, solutions.Day04Solutions)

	cobra.CheckErr(rootCmd.Execute())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".advent-of-code-2020" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".advent-of-code-2020")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
