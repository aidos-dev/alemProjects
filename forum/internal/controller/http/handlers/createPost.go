package handlers

import (
	"errors"
	"forum/internal/infrastructure/repository"
	"forum/internal/model"
	"forum/internal/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CreatePost struct {
	serv service.Post
}

func CreateCreatePostHandler(serv service.Post) *CreatePost {
	return &CreatePost{
		serv: serv,
	}
}

func (cp CreatePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(r)
	user := r.Context().Value("authorizedUser").(*model.User)
	if user.ID == 0 {
		errorPage(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, w, r)
		return
	}
	if r.Method == http.MethodGet { // 14 если приходит метод гет, рендерим страницу
		t, err := template.ParseFiles("./templates/createPostAuth.html")
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
		err = t.Execute(w, user.Username)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
		return
	}
	if r.Method == http.MethodPost { // 14 если пост - создаем пост
		err := r.ParseForm()
		if err != nil {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}

		postInfo := r.PostForm // 14 это мапа у который ключ стринга, а значение -  массив стрингов
		for key := range postInfo {
			if !contains([]string{"category", "text", "heading"}, key) { // 14 проверяем все поля формы, чтобы они были правильные
				errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
				return
			}
		}

		values := r.Form["category"]
		for _, value := range values {
			if !contains([]string{"1", "2", "3", "4", "5", "6"}, value) {
				errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
				return
			}
		}

		for i, category := range values {
			for _, nextCategory := range values[i+1:] {
				if category == nextCategory {
					errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
					return
				}
			}
		}

		post := model.PostRepresentation{ // 14 создается сам пост
			Heading: postInfo["heading"][0], //[0] так как одно поле
			Text:    postInfo["text"][0],
		}

		// empty post heading or emty post text check
		if strings.ReplaceAll(post.Heading, " ", "") == "" || len(post.Heading) < 1 || len(post.Heading) > 200 || strings.ReplaceAll(post.Text, " ", "") == "" || len(post.Text) < 6 || len(post.Text) > 1500 {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}

		postId, err := cp.serv.CreatePost(post.Heading, post.Text, user.ID) // 14 сохраняется в БД
		if errors.Is(err, repository.ErrInvalidPost) {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		} else if err != nil {
			log.SetFlags(log.Llongfile)
			log.Println("HERE", err)
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}

		categories := postInfo["category"]
		for _, categoryId := range categories { // 14 проходимся по всем категориям сохраняем в БД интовым значением 1-6
			categoryIdUint, _ := strconv.ParseUint(categoryId, 10, 32)
			_, err := cp.serv.AddCategoryToPost(uint(categoryIdUint), postId)
			if err != nil {
				errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
				return
			}
		}

		postIdString := strconv.FormatUint(uint64(postId), 10)
		http.Redirect(w, r, "/post?id="+postIdString, http.StatusSeeOther) // 14 редиректим юзера на страницу только что созданного поста
		return
	} else {
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w, r) // 14 если пришел другой метод, возвращаем ошибку
		return
	}
}
