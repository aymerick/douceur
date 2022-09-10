/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/lab42/douceur/parser"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a CSS file and display result",
	Run: func(cmd *cobra.Command, args []string) {
		input := readFile(cmd.Flag("file").Value.String())

		stylesheet, err := parser.Parse(string(input))
		cobra.CheckErr(err)

		fmt.Println(stylesheet.String())
	},
}

func init() {
	parseCmd.Flags().StringP("file", "f", "", "File")
	parseCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(parseCmd)
}
