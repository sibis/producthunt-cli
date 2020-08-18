/*
copyright © 2020 Sibi <sibi.kabil@gmail.com>

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
	"net/url"
	"os/exec"

	"github.com/spf13/cobra"
)

// signinCmd represents the signin command
var signinCmd = &cobra.Command{
	Use:   "signin",
	Short: "Authorize the application from your account",
	Long:  `Helps to generate the code by providing the access, we promise not to pose without your consent`,
	Run: func(cmd *cobra.Command, args []string) {
		client_id := "K0G_mZKnQkvmBPTHXC7bKAXgJZlLzgA0TePqFpn2yJU"
		q := url.Values{}
		q.Set("client_id", client_id)
		q.Set("scope", "public")
		q.Set("redirect_uri", fmt.Sprintf("https://producthuntcli.netlify.app/"))
		q.Set("response_type", "code")
		url := fmt.Sprintf("https://api.producthunt.com/v2/oauth/authorize?%s", q.Encode())
		err := exec.Command("open", url).Start()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(signinCmd)
}
