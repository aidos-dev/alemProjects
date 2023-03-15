package handlers

import (
	"database/sql"
	"errors"
	"forum/internal/model"
	"net/http"
	"strconv"
)

type Reaction interface {
	React(postOrComment string, userId uint, postId uint, positive bool) (uint, error)
}

type React struct {
	reactionServ Reaction
	postServ     PostServices
}

func CreateReactHandler(reactionServ Reaction, postServ PostServices) *React {
	return &React{
		reactionServ: reactionServ,
		postServ:     postServ,
	}
}

func (re React) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("authorizedUser").(*model.User)
	if user.ID == 0 {
		errorPage(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, w, r)
		return
	}

	if r.Method != http.MethodPost {
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	formInfo := r.PostForm // 14 это мапа у который ключ стринга, а значение -  массив стрингов
	for key := range formInfo {
		if !contains([]string{"reactTo", "postId", "positive", "commentId"}, key) { // 14 проверяем все поля формы, чтобы они были правильные
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}
	}

	postOrComment, ok := r.PostForm["reactTo"]
	if !ok {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	postIdString, ok := r.Form["postId"]
	if !ok {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	postId, err := strconv.Atoi(postIdString[0])
	if err != nil {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	_, err = re.postServ.GetPostById(uint(postId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//	fmt.Println("asdasd")
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
		return
	}

	reaction, ok := r.Form["positive"]
	if !ok {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	positive, err := strconv.ParseBool(reaction[0])
	if err != nil {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	if postOrComment[0] == "post" {
		postId, err := strconv.Atoi(postIdString[0])
		if err != nil {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}

		postIdUint := uint(postId)

		_, err = re.reactionServ.React(postOrComment[0], user.ID, postIdUint, positive) // 19 ставится реакция (заходим в reaction.go 24 строка)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
	} else if postOrComment[0] == "comment" {
		commentIdString, ok := r.Form["commentId"]
		if !ok {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}

		commentId, err := strconv.Atoi(commentIdString[0])
		if err != nil {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}

		commentIdUint := uint(commentId)

		_, err = re.reactionServ.React(postOrComment[0], user.ID, commentIdUint, positive)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
	} else {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	http.Redirect(w, r, "/post?id="+postIdString[0], http.StatusSeeOther)
	return
}
