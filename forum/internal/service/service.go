package service

type Repository interface {
	RegisterUser
	LoginUser
	SessionCreator
	SessionChecker
	Post
	Comment
	React
}

type Service struct {
	RegisterUserService  RegisterUserService
	LoginUserService     LoginUserService
	SessionCreateService SessionCreateService
	SessionCheckService  SessionCheckService
	PostService          PostService
	CommentService       CommentService
	ReactService         ReactService
}

func NewService(repo Repository) *Service { //  7,2 сборка реализаций всех сервисных методов бизнес логики
	return &Service{
		RegisterUserService:  *NewRegisterUserService(repo),
		LoginUserService:     *NewLoginUserService(repo),
		SessionCreateService: *NewSessionCreateService(repo),
		SessionCheckService:  *NewSessionCheckService(repo),
		PostService:          *NewPostService(repo),
		CommentService:       *NewCommentService(repo),
		ReactService:         *NewReacttService(repo),
	}
}
