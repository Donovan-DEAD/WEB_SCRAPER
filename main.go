package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	checksuccessive "github.com/Donovan-DEAD/WEB_SCRAPER/packages/utils/checkSuccessive"
)

func main() {
	stdinReader := bufio.NewReader(os.Stdin)

	chanel := make(chan []string)
	website := "https://www.w3schools.com/"

	linksInWebsite := sync.Map{}
	linksToCheck := []string{website}

	count := 1

	for {
		decision := ""
		fmt.Println("\n\nYou want to continue scanning the web site for more links? y/n\n\n*Consider each loop increases several times the time it takes to complete.")
		stdinReader.ReadString('\n')
		fmt.Scan(decision)

		if decision == "N" || decision == "n" {
			break
		}

		fmt.Println("Loop number: ", count)
		count++
		temporalSlice := linksToCheck
		linksToCheck = []string{}

		for index, path := range temporalSlice {
			go checksuccessive.Checksuccessive(path, chanel, &linksInWebsite)
			time.Sleep(time.Millisecond * 10)

			if index%150 == 0 || index == len(temporalSlice) {
				checksuccessive.WaitForPoolEntries(&linksToCheck, chanel)
			}
		}
		fmt.Println("\n\nElements in the map:")
		countForRange := 1
		linksInWebsite.Range(func(key, value any) bool {
			fmt.Println(value, "\t", countForRange, "\t", key)
			countForRange++
			return true
		})

		if len(linksToCheck) == 0 {
			break
		}

	}

	close(chanel)

	countForRange := 1
	linksInWebsite.Range(func(key, value any) bool {
		fmt.Println(value, "\t", countForRange, "\t", key)
		countForRange++
		return true
	})
}
