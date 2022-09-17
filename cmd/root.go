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
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "douceur",
	Short: "A simple CSS parser and inliner in Golang.",
	Long: `Parser is vaguely inspired by CSS Syntax Module Level 3 and corresponding JS parser. 
Inliner only parses CSS defined in HTML document, it DOES NOT fetch external stylesheets (for now). 
Inliner inserts additional attributes when possible.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func readFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to open file: ", filePath, err)
		os.Exit(1)
	}

	return file
}
