package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/mgutz/ansi"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

type PostJSON struct {
	Data struct {
		Post struct {
			Name          string  `json:"name"`
			Description   string  `json:"description"`
			ID            string  `json:"id"`
			Tagline       string  `json:"tagline"`
			Website       string  `json:"website"`
			ReviewRatings float64 `json:"reviewsRating"`
			Makers        []struct {
				Name            string `json:"name"`
				TwitterUsername string `json:"twitterUsername"`
				Headline        string `json:"headline"`
				WebsiteURL      string `json:"websiteUrl"`
			} `json:"makers"`
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
		} `json:"post"`
	} `json:"data"`
}

func Readtoken() string {
	path, _ := os.LookupEnv("HOME")
	dat, err := ioutil.ReadFile(path + "/.producthunt")
	if err != nil {
		panic(err)
	}
	result := string(dat)
	return result
}

func ViewProduct(args []string) {
	fmt.Println("\n")

	heading := ansi.ColorFunc("green")
	productName := ansi.ColorFunc("cyan+b")
	linkColor := ansi.ColorFunc("blue")
	size, _ := terminal.Width()

	jsonData := map[string]string{
		"query": `
                {
                    post (id:` + string(args[0]) + `) {
                        name
                        website
                        tagline
                        description
                        isVoted
                        makers {
                            name
                            twitterUsername
      headline
      websiteUrl
                        }
                        reviewsRating
                        productLinks {
                            type
                            url
                        }
                        topics {
                            edges {
                                node {
                                    name
                                }
                            }
                        }
                    }
                }
            `}
	data := MakeReq(jsonData)

	var p PostJSON
	_ = json.Unmarshal([]byte(data), &p)
	fmt.Println("\n")
	headingmsg := heading(p.Data.Post.Name)
	fmt.Printf(fmt.Sprintf("%%-%ds", int(size)/2), fmt.Sprintf(fmt.Sprintf("%%%ds", int(size)/2), headingmsg))
	fmt.Println("\n")
	fmt.Println(productName("Tagline : "), p.Data.Post.Tagline, "\n")
	fmt.Println(productName("Description : "), p.Data.Post.Description, "\n")

	fmt.Println(productName("Maker(s) : "))
	for midx, maker := range p.Data.Post.Makers {
		fmt.Println("\t", fmt.Sprintf("👤 (%d) %s", midx+1, maker.Name))
		if len(maker.TwitterUsername) > 0 {
			fmt.Println("\t", fmt.Sprintf("🐦 %s", maker.TwitterUsername))
		}
		if len(maker.Headline) > 0 {
			fmt.Println("\t", fmt.Sprintf("📰 %s", maker.Headline))
		}
		if len(maker.WebsiteURL) > 0 {
			fmt.Println("\t", fmt.Sprintf("ℹ️  %s", linkColor(maker.WebsiteURL)))
		}
		fmt.Println("\n")
	}

	productCategory := ""
	for idx, category := range p.Data.Post.Topics.Edges {
		if idx != 0 {
			productCategory += ", "
		}
		productCategory += category.Node.Name
	}
	fmt.Println(productName("Categories : "), productCategory, "\n")

	for _, link := range p.Data.Post.ProductLinks {
		fmt.Println(productName(fmt.Sprintf("%s", link.Type))+" : ", fmt.Sprintf("%s", linkColor(link.URL)))
	}
	fmt.Println("\n")
	fmt.Println(productName("Avg. reviews rating : "), p.Data.Post.ReviewRatings, "\n")

}

func ChooseProduct(trendingNames []string, trendingNamesMap map[string]string) {
	prompt := promptui.Select{
		Label: "Select the product you read the details: ",
		Items: trendingNames,
	}

	_, result, err1 := prompt.Run()

	if err1 != nil {
		fmt.Printf("Prompt failed %v\n", err1)
		return
	}

	var argName []string
	for key, value := range trendingNamesMap {
		if value == result {
			argName = append(argName, key)
			ViewProduct(argName)
		}
	}
	ChooseProduct(trendingNames, trendingNamesMap)
}

func MakeReq(jsonData map[string]string) []byte {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond) // Build our new spinner
	s.Start()

	jsonValue, _ := json.Marshal(jsonData)

	url := "https://api.producthunt.com/v2/api/graphql"

	token := Readtoken()
	var bearer = "Bearer " + token
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	s.Stop()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return data
}
