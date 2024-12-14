package searchLinks

import (
	"regexp"
	"strings"
)

func FilterPaths(linksFounded []string) (linksFiltered []string) {

	fileExtensionsToAvoid := []string{
		".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt", ".rtf", ".zip",
		".rar", ".7z", ".tar", ".gz", ".bz2", ".jpg", ".jpeg", ".png", ".gif", ".bmp",
		".svg", ".mp3", ".wav", ".flac", ".aac", ".ogg", ".mp4", ".avi", ".mkv", ".mov",
		".wmv", ".exe", ".dmg", ".iso", ".apk", ".msi", ".epub", ".mobi", ".azw", ".csv",
		".json", ".xml", ".ico", ".js", ".css", ".woff2", "javascript:void(0)"}

	for _, value := range linksFounded {

		eliminate := false

		for _, suffix := range fileExtensionsToAvoid {

			if strings.Contains(value, suffix) {
				eliminate = true
				break
			} else {
				continue
			}
		}

		if strings.HasPrefix(value, "https:") || strings.HasPrefix(value, "http:") {
			eliminate = true
		}

		if !eliminate {
			linksFiltered = append(linksFiltered, value)
		}

	}
	return linksFiltered
}

func QuitHtmlProperty(linksFiltered []string) (linksFinal []string) {
	regEx := regexp.MustCompile("\"[^\"]*")

	for _, value := range linksFiltered {
		linksFinal = append(linksFinal, regEx.FindString(value)[1:])
	}

	return linksFinal
}

func SearchForlinksR(page string) (linksInPage []string) {
	patternRelative := "\".*\""
	regexToSearch := []string{"href=", "src=", "action=", "cite=", "background=", "data=", "poster="}

	for _, value := range regexToSearch {
		regEx := regexp.MustCompile(value + patternRelative)
		linksInPage = append(linksInPage, regEx.FindAllString(page, -1)...)
	}

	linksInPage = QuitHtmlProperty(linksInPage)
	linksInPage = FilterPaths(linksInPage)

	return linksInPage
}
