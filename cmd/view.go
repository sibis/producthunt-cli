/*
 Copyright © 2020 Sibi <sibi.kabil@gmail.com>

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
	"../utils"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "To view the prodcut information",
	Long:  `A detailed information related to the product and its makers.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ViewProduct(args)
	},
}

func init() {
	// Product flag needs to be passed to execute `ph view` command
	viewCmd.Flags().StringP("product", "p", "", "Product ID (required)")
	rootCmd.AddCommand(viewCmd)
}
