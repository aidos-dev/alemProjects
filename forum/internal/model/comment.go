package model

type Comment struct {
	Text   string
	PostId uint
	UserUd uint
}

type CommentRepresentation struct {
	CommentId      uint
	Text           string
	AmountLikes    int
	AmountDisLikes int
	Username       string
	UserId         uint
	PostId         uint
}
