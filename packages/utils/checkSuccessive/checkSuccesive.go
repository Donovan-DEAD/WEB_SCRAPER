package checksuccessive

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Donovan-DEAD/Web_Scraper/packages/models/links"
	"github.com/Donovan-DEAD/Web_Scraper/packages/utils/searchLinks"
)

func checksuccessive(linkToCheck string, channel chan links.Link) {
	resp, err := http.Get(linkToCheck)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	linksInPage := searchLinks.SearchForlinks(string(body))

	for _, value := range linksInPage {
		if strings.HasSuffix(linkToCheck, "/") && strings.HasPrefix(value, "/") {

			go checksuccessive(linkToCheck+value[1:], channel)

		} else {

			go checksuccessive(linkToCheck+value, channel)

		}
	}

	channel <- links.CreateNewLink(linkToCheck, resp.StatusCode)

	defer resp.Body.Close()
}
