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
		".json", ".xml", ".ico", ".js", ".css"}

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

		if !eliminate {
			linksFiltered = append(linksFiltered, value)
		}

	}
	return linksFiltered
}

func QuitHtmlProperty(linksFiltered []string) (linksFinal []string) {
	regEx := regexp.MustCompile("/[^\"]*")

	for _, value := range linksFiltered {
		linksFinal = append(linksFinal, regEx.FindString(value))
	}

	return linksFinal
}

func SearchForlinks(page string) (linksInPage []string) {

	regexToSearch := []string{"href=\"/[^\"]*\"", "src=\"/[^\"]*\"", "action=\"/[^\"]*\"", "cite=\"/[^\"]*\"", "background=\"/[^\"]*\"", "data=\"/[^\"]*\"", "poster=\"/[^\"]*\""}

	for _, value := range regexToSearch {
		regEx := regexp.MustCompile(value)
		linksInPage = append(linksInPage, regEx.FindAllString(page, -1)...)
	}

	linksInPage = FilterPaths(linksInPage)
	linksInPage = QuitHtmlProperty(linksInPage)

	return linksInPage
}
