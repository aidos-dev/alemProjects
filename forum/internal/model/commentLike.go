package model

type CommentLike struct {
	CommentLikeId uint
	UserId        uint
	PostId        uint
	Positive      bool
}
