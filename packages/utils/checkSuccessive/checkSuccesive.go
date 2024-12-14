package checksuccessive

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/Donovan-DEAD/WEB_SCRAPER/packages/utils/searchLinks"
)

func Checksuccessive(linkToCheck string, channel chan []string, linkRegister *sync.Map) {

	//Starts the http request
	resp, err := http.Get(linkToCheck)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//If the status of the link is dead return it before it reads the body and cause an error
	if resp.StatusCode != http.StatusOK {
		linkRegister.Store(linkToCheck, resp.StatusCode)
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
	linkRegister.Store(linkToCheck, resp.StatusCode)

	linksInPageFiltered := []string{}

	//For each link start a new go routine
	for _, value := range linksInPage {

		urlStructure, _ := url.Parse(linkToCheck)
		//If the value for some reason has no characters is continue to not cause a runtime error
		if len(value) == 0 {
			continue
		}
		//Check if the path is relative to the root of the page or relative to the actual page
		if string(value[0]) == "/" {
			urlStructure.Path = value
		} else {
			regEx := regexp.MustCompile("/[^/]*")

			pathDestructured := regEx.FindAllString(urlStructure.Path, -1)
			pathDestructured[len(pathDestructured)-1] = value

			urlStructure.Path = strings.Join(pathDestructured, "")
		}

		if _, present := linkRegister.Load(urlStructure.String()); present {
			continue
		} else {
			linksInPageFiltered = append(linksInPageFiltered, urlStructure.String())
		}
	}

	if len(linksInPageFiltered) == 0 {
		return
	}

	channel <- linksInPageFiltered
}

func WaitForPoolEntries(linksToCheck *[]string, chanel chan []string) {

	breakFor := false

	for !breakFor {
		select {

		case value := <-chanel:
			(*linksToCheck) = append(*linksToCheck, value...)

		case <-time.After(time.Second * 1):
			breakFor = true

		}
	}
}
