package model

type PostLike struct {
	PostLikeId uint
	UserId     uint
	PostId     uint
	Positive   bool
}
