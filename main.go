package main

import (
	"fmt"
	"time"

	"github.com/Donovan-DEAD/Web_Scraper/packages/models/links"
	checksuccessive "github.com/Donovan-DEAD/Web_Scraper/packages/utils/checkSuccessive"
)

func main() {
	chanel := make(chan links.Link)

	linksInWebsite := map[string]int{}

	website := "https://www.google.com/"

	go checksuccessive.Checksuccessive(website, chanel)

	breakFor := false
	for !breakFor {
		select {

		case value := <-chanel:

			if linksInWebsite[value.Path] == 0 {
				linksInWebsite[value.Path] = value.StatusCode
			}

		case <-time.After(time.Second * 5):
			breakFor = true
		}
	}

	for key, value := range linksInWebsite {
		fmt.Println(key, "\t", value)
	}
}
