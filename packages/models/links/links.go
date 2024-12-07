package links

type Link struct {
	path       string
	statusCode int
	dead       bool
}

func CreateNewLink(paths string, status int) (newLink Link) {
	newLink.path = paths
	newLink.statusCode = status

	if newLink.statusCode < 400 {
		newLink.dead = false
	} else {
		newLink.dead = true
	}

	return newLink
}
