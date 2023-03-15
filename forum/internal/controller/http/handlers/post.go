package handlers

import (
	"database/sql"
	"errors"
	"forum/internal/model"
	"html/template"
	"net/http"
	"strconv"
)

// это все страница одного поста

type PostService interface {
	GetPostById(postId uint) (*model.PostRepresentation, error)
}

type CommentService interface {
	GetAllCommentsByPostId(postId uint) ([]model.CommentRepresentation, error)
}

type Post struct {
	postSvc    PostService
	commentSvc CommentService
}

func CreatePostHandler(postSvc PostService, commentSvc CommentService) *Post {
	return &Post{
		postSvc:    postSvc,
		commentSvc: commentSvc,
	}
}

func (p *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w, r)
		return
	}
	user := r.Context().Value("authorizedUser").(*model.User)
	postIdSting := r.URL.Query().Get("id") // 18 получает айди поста

	postId64, err := strconv.ParseUint(postIdSting, 10, 32) // 18,1 переводит в число
	if err != nil {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w, r)
		return
	}

	postId := uint(postId64)

	post, err := p.postSvc.GetPostById(postId) // 18,1 получает пост
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // 18,1 если поста нет вернется ошибка sql.ErrNoRows
			errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w, r)
			return
		}
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
		return
	}

	comments, err := p.commentSvc.GetAllCommentsByPostId(postId)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
		return
	}

	// 18,1 если авторизованный
	info := struct {
		User     *model.User
		Post     *model.PostRepresentation
		Auth     bool
		Comments []model.CommentRepresentation
	}{
		User:     user,
		Post:     post,
		Auth:     true,
		Comments: comments,
	}
	// 18,1 если не авторизованный
	if user.ID == 0 {
		info1 := struct {
			Post     *model.PostRepresentation
			Auth     bool
			Comments []model.CommentRepresentation
		}{
			Post:     post,
			Auth:     false,
			Comments: comments,
		}
		t, err := template.New("post.html").Funcs(template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		}).ParseFiles("./templates/post.html")
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
		err = t.Execute(w, info1)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
		return
	}
	// function inside template если авторизованный
	t, err := template.New("post.html").Funcs(template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}).ParseFiles("./templates/post.html")
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
		return
	}
	err = t.Execute(w, info)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
		return
	}
}
