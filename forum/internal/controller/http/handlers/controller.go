package handlers

import (
	"forum/internal/service"
	"net/http"
)

type Service interface {
	Registration
	Authorization
	Auth
	Leaving
	PostService
	CommentService
	Reaction
}

type Controller struct {
	index
	Logout
	SignUp
	SingIn
	Middleware
	CreatePost
	Post
	CreateComment
	React
	Filter
}

func NewContoller(serv *service.Service) *Controller { // 7,3 сборка всех хэндлеров
	return &Controller{
		index:         *createIndexHandler(&serv.PostService),
		Logout:        *CreateLogoutHandler(&serv.SessionCheckService),
		SignUp:        *CreateSignUpHandler(&serv.RegisterUserService),
		SingIn:        *CreateSignInHandler(&serv.LoginUserService, &serv.SessionCreateService),
		Middleware:    *CreateMiddleware(&serv.SessionCheckService),
		CreatePost:    *CreateCreatePostHandler(&serv.PostService),
		Post:          *CreatePostHandler(&serv.PostService, &serv.CommentService),
		CreateComment: *CreateCommentHandler(&serv.CommentService, &serv.PostService),
		React:         *CreateReactHandler(&serv.ReactService, &serv.PostService),
		Filter:        *CreateFilterHandler(&serv.PostService),
	}
}

func (c *Controller) InitRouter() *http.ServeMux { // 10,1 сборка роутера, где прописываются все пути и какой хэндлер используется и сразу оборачиывается мидлвеаром
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.Handle("/likeup", c.LoggerMiddlerware(c.AuthMiddleware(c.React))) // 10,1 логгермидлвеар добавился для того, чтобы логировать все происхлдящее в консоль
	mux.Handle("/likedown", c.LoggerMiddlerware(c.AuthMiddleware(c.React)))
	mux.Handle("/", c.LoggerMiddlerware(c.AuthMiddleware(&c.index)))
	mux.Handle("/logout", c.LoggerMiddlerware(c.AuthMiddleware(c.Logout))) // 17 !!!! добавили AuthMiddleware чтобы неавторизованный пользователь не мог попасть на страницу логаута
	mux.Handle("/signup", c.LoggerMiddlerware(c.AuthMiddleware(&c.SignUp)))
	mux.Handle("/signin", c.LoggerMiddlerware(c.AuthMiddleware(&c.SingIn)))
	mux.Handle("/create-post", c.LoggerMiddlerware(c.AuthMiddleware(c.CreatePost)))
	mux.Handle("/post", c.LoggerMiddlerware(c.AuthMiddleware(&c.Post)))

	mux.Handle("/create-comment", c.LoggerMiddlerware(c.AuthMiddleware(c.CreateComment)))

	mux.Handle("/filter", c.LoggerMiddlerware(c.AuthMiddleware(&c.Filter)))

	return mux
}
