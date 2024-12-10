package checksuccessive

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Donovan-DEAD/Web_Scraper/packages/models/links"
	"github.com/Donovan-DEAD/Web_Scraper/packages/utils/searchLinks"
)

func Checksuccessive(linkToCheck string, channel chan links.Link, linkRegister *map[string]int) {
	if (*linkRegister)[linkToCheck] != 0 {
		return
	}

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

	linksInPage := searchLinks.SearchForlinksR(string(body))

	for _, value := range linksInPage {
		if strings.HasSuffix(linkToCheck, "/") && strings.HasPrefix(value, "/") {

			go Checksuccessive(linkToCheck+value[1:], channel, linkRegister)

		} else {

			go Checksuccessive(linkToCheck+value, channel, linkRegister)

		}
	}

	channel <- links.CreateNewLink(linkToCheck, resp.StatusCode)

	defer resp.Body.Close()
}
