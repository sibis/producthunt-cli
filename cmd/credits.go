/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

var creditsCmd = &cobra.Command{
	Use:   "credits",
	Short: "A short information about ther maker of this tool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n")
		fmt.Println(`

		**********************************************************************************************************
		*                                                                                                        *
		*   Thanks for using ph! 💫                                                                              *
		*                                                                                                        *
		*   I'm Sibi from India, backend developer works mostly on Go and Python.                                *
		*   Porfolio: https://sibis.dev                                                                          *
		*                                                                                                        *
		*   If you like this app. Please consider upvoting it on Product Hunt and star it on Github.             *
		*   Entire repository is open-source.                                                                    *
		*   Source code: https://github.com/sibis/producthunt-cli                                                *
		*                                                                                                        *
		*   I would like to hear your feedback on this tool. My DMs are always open and love to hear from you.   *
		*   Twitter: https://twitter.com/_sibis                                                                  *
		*                                                                                                        *
		**********************************************************************************************************

		
		`)
	},
}

func init() {
	rootCmd.AddCommand(creditsCmd)
}
