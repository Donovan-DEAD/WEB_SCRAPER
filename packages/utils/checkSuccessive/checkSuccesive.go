package checksuccessive

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/Donovan-DEAD/Web_Scraper/packages/models/links"
	"github.com/Donovan-DEAD/Web_Scraper/packages/utils/searchLinks"
)

func Checksuccessive(linkToCheck string, channel chan links.Link, linkRegister *map[string]int) {

	//Check if the main program has already that link
	if (*linkRegister)[linkToCheck] != 0 {
		return
	}

	//Starts the http request
	resp, err := http.Get(linkToCheck)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//If the status of the link is dead return it before it reads the body and cause an error
	if resp.StatusCode != http.StatusOK {
		channel <- links.CreateNewLink(linkToCheck, resp.StatusCode)
		return
	}

	//Read the body so the program can search for links
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	//Read all the links in the html
	linksInPage := searchLinks.SearchForlinksR(string(body))

	//Send the result to the main program so they can register the link and its status code so they wont be repeated
	channel <- links.CreateNewLink(linkToCheck, resp.StatusCode)

	//For each link start a new go routine
	for _, value := range linksInPage {
		for {
			numGoRoutines := runtime.NumGoroutine()

			//Check if the number of go routines dont increase a lot so the computer can handle with its resources
			if numGoRoutines <= 130 {
				break
			}

			time.Sleep(10 * time.Millisecond)

		}

		urlStructure, _ := url.Parse(linkToCheck)

		//Check if the path is relative to the root of the page or relative to the actual page
		if string(value[0]) == "/" {
			urlStructure.Path = value
		} else {
			urlStructure.Path += value
		}

		go Checksuccessive(urlStructure.String(), channel, linkRegister)
	}

}
