package links

type Link struct {
	Path       string
	StatusCode int
}

func CreateNewLink(paths string, status int) (newLink Link) {
	newLink.Path = paths
	newLink.StatusCode = status

	return newLink
}
