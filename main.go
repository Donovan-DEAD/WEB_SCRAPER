package main

import (
	"fmt"
	"sync"
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

		case <-time.After(time.Second * 30):

			breakFor = true

		}
	}

	var wg sync.WaitGroup
	wg.Wait()
	close(chanel)

	count := 0
	for key, value := range linksInWebsite {
		fmt.Println(value, "\t", count, "\t", key)
		count++
	}
}
