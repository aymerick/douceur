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

	"github.com/lab42/douceur/inliner"
	"github.com/spf13/cobra"
)

// inlineCmd represents the inline command
var inlineCmd = &cobra.Command{
	Use:   "inline",
	Short: "Inline CSS in an HTML document and display result",
	Run: func(cmd *cobra.Command, args []string) {
		input := readFile(cmd.Flag("file").Value.String())

		output, err := inliner.Inline(string(input))
		cobra.CheckErr(err)

		fmt.Println(output)
	},
}

func init() {
	inlineCmd.Flags().StringP("file", "f", "", "File")
	inlineCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(inlineCmd)
}
