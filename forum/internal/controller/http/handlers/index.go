package handlers

import (
	"forum/internal/model"
	"html/template"
	"net/http"
)

type IndexInterface interface {
	GetAllPosts() ([]model.PostRepresentation, error)
}

type index struct {
	serv IndexInterface
}

func createIndexHandler(serv IndexInterface) *index {
	return &index{
		serv: serv,
	}
}

func (i index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w, r)
		return
	}

	if r.Method != http.MethodGet { // 17 проверка на метод
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w, r)
		return
	}

	user := r.Context().Value("authorizedUser").(*model.User)
	allposts, err := i.serv.GetAllPosts()
	if err != nil {
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
		return
	}
	info := struct {
		User          *model.User
		Posts         []model.PostRepresentation
		HeadingFilter string
	}{
		User:          user,
		Posts:         allposts,
		HeadingFilter: "Latest Posts",
	}
	if user.ID == 0 {
		t, err := template.New("index.html").Funcs(template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		}).ParseFiles("./templates/index.html")
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
		err = t.Execute(w, info)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
		return
	}
	t, err := template.New("indexAuthorized.html").Funcs(template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}).ParseFiles("./templates/indexAuthorized.html")
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
