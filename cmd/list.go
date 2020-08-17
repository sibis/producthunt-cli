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
	"encoding/json"
	"fmt"

	"../utils"
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

type PostsJson struct {
	Data struct {
		Posts struct {
			Edges []struct {
				Node struct {
					Name         string `json:"name"`
					VotesCount   int    `json:"votesCount"`
					Description  string `json:"description"`
					ID           string `json:"id"`
					Website      string `json:"website"`
					Tagline      string `json:"tagline"`
					ProductLinks []struct {
						URL  string `json:"url"`
						Type string `json:"type"`
					} `json:"productLinks"`
					Topics struct {
						Edges []struct {
							Node struct {
								Name string `json:"name"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"topics"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"posts"`
	} `json:"data"`
}

// signinCmd represents the signin command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Fetches the list of products trnding",
	Long:  `current trending products`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonData := map[string]string{
			"query": `
                    {
                    posts {
                        edges {
                            node {
                                id
                                name
                                votesCount
                                tagline
                                productLinks {
                                    type
                                    url
                                }
                                website
                                topics {
                                    edges {
                                        node {
                                            name
                                        }
                                    }
                                }

                            }
                        }
                    }
            }
        `,
		}
		data := utils.MakeReq(jsonData)

		fmt.Println("\n")
		heading := ansi.ColorFunc("green")
		productName := ansi.ColorFunc("cyan+b")
		categoryColor := ansi.ColorFunc("magenta")
		linkColor := ansi.ColorFunc("blue")
		headingmsg := heading("Trending products")
		size, _ := terminal.Width()
		fmt.Printf(fmt.Sprintf("%%-%ds", int(size)/2), fmt.Sprintf(fmt.Sprintf("%%%ds", int(size)/2), headingmsg))
		fmt.Println("\n")

		var p PostsJson
		_ = json.Unmarshal([]byte(data), &p)

		trendingNamesMap := make(map[string]string)
		var trendingNames []string
		for _, product := range p.Data.Posts.Edges {

			trendingNamesMap[product.Node.ID] = product.Node.Name
			trendingNames = append(trendingNames, product.Node.Name)

			productCategory := " / [ "
			for idx, category := range product.Node.Topics.Edges {
				if idx != 0 {
					productCategory += ", "
				}
				productCategory += fmt.Sprintf("%s", category.Node.Name)
			}
			productCategory += " ]"
			fmt.Print("👉 " + productName(product.Node.Name) + " (↑" + string(product.Node.VotesCount) + ") " + categoryColor(productCategory) + "\n")
			fmt.Println(product.Node.Tagline)

			for _, link := range product.Node.ProductLinks {
				fmt.Println(link.Type+" - ", linkColor(link.URL))
			}

			fmt.Println("\n")

		}

		utils.ChooseProduct(trendingNames, trendingNamesMap)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
