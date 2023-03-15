package handlers

import (
	"database/sql"
	"errors"
	"forum/internal/model"
	"net/http"
	"strconv"
	"strings"
)

type Comment interface { // 11 интефейс который описывает метод создания коммента
	CreateComment(userId, postId uint, text string) (uint, error)
}

type PostServices interface {
	GetPostById(postId uint) (*model.PostRepresentation, error)
}

type CreateComment struct {
	commentServ Comment
	postServ    PostServices
}

func CreateCommentHandler(commentServ Comment, postServ PostServices) *CreateComment { // 11,1 принимает интерфейс, который работает благодаря (открыть controller.go 42 строка)+(открыть service.go 19 строка)
	return &CreateComment{
		commentServ: commentServ,
		postServ:    postServ,
	}
}

func (cc CreateComment) ServeHTTP(w http.ResponseWriter, r *http.Request) { // 11,1 реализация метода servHttp интерфейса http.Handler
	curUser := r.Context().Value("authorizedUser").(*model.User) // 11,1 здесь точка value возврзает значени по ключу как мапа (ключ - "authorizedUser") (открыть middleware.go 28, там дальше 11,1)
	if curUser.ID == 0 {
		errorPage(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, w, r)
		return
	}

	if r.Method != http.MethodPost {
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w, r)
		return
	}

	err := r.ParseForm() // 12 парсится форма
	if err != nil || len(r.Form) < 3 {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	commentInfo := r.PostForm

	if _, ok := commentInfo["comment"]; !ok {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	if _, ok := commentInfo["postId"]; !ok {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	// if _, ok := commentInfo["userId"]; !ok {
	// 	errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
	// 	return
	// }

	if len(commentInfo["postId"]) < 1 || len(commentInfo["comment"]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	postIdString := commentInfo["postId"][0]
	postId64, err := strconv.ParseUint(postIdString, 10, 32)
	if err != nil {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}
	postId := uint(postId64)

	comment := model.CommentRepresentation{
		PostId: postId,
		UserId: curUser.ID,
		Text:   commentInfo["comment"][0],
	}

	// empty comment text check
	if strings.ReplaceAll(comment.Text, " ", "") == "" || len(comment.Text) < 1 || len(comment.Text) > 1500 || comment.UserId <= 0 || comment.PostId <= 0 {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	_, err = cc.postServ.GetPostById(postId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//	fmt.Println("asdasd")
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
		return
	}

	_, err = cc.commentServ.CreateComment(comment.UserId, comment.PostId, comment.Text) // 12,1 создается коммент (перейти service - comment.go)
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
		return
	}

	http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther)
}
