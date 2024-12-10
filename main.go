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

	website := ""

	go checksuccessive.Checksuccessive(website, chanel, &linksInWebsite)

	breakFor := false
	for !breakFor {
		select {

		case value := <-chanel:

			linksInWebsite[value.Path] = value.StatusCode

		case <-time.After(time.Second * 1):

			breakFor = true

		}
	}

	for key, value := range linksInWebsite {
		fmt.Println(key, "\t", value)
	}
}
