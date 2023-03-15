package model

type Post struct {
	PostId   uint
	Heading  string
	Text     string
	UserId   uint
	Category string
}

type PostRepresentation struct {
	PostId         uint
	Heading        string
	Text           string
	AmountLikes    int
	AmountDisLikes int
	AmountComments uint
	Username       string
	Categories     []string
}
