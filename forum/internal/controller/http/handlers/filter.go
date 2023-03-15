package handlers

import (
	"forum/internal/model"
	"html/template"
	"net/http"
	"strings"
)

type FilterInterface interface { // 16 методы которые используются в хэндлере
	FilterAllPosts(filterBy string) ([]model.PostRepresentation, error)
	PersonalFilter(filterBy string, userId uint) ([]model.PostRepresentation, error)
}

type Filter struct {
	serv FilterInterface
}

func CreateFilterHandler(serv FilterInterface) *Filter { // 16,1 принимает интерфейс отдает структуру
	return &Filter{
		serv: serv,
	}
}

func (f Filter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w, r)
		return
	}

	if r.URL.Path != "/filter" {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w, r)
		return
	}

	var err error
	user := r.Context().Value("authorizedUser").(*model.User)
	if user.ID == 0 { // 16,1 если не авторизованный
		err = r.ParseForm() // 16,1 считывает html и парсит
		if err != nil {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}

		if _, ok := r.Form["filter_by"]; !ok {
			errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w, r)
			return
		}

		value := r.Form["filter_by"]
		if !contains([]string{"oldest", "recent", "most_disliked", "most_liked", "discussions", "questions", "ideas", "articles", "events", "issues"}, value[0]) {
			errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w, r)
			return
		}

		filterBy := r.FormValue("filter_by") // 16,1 берет значение из query, а query это все что после вопроситеоьного знака

		var filteredPosts []model.PostRepresentation

		filteredPosts, err = f.serv.FilterAllPosts((filterBy)) // 16,1 получаем все отфилтрованные посты
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}

		info := struct { // 16,1 эта инфа будет отрисовываться в html
			User          *model.User
			Posts         []model.PostRepresentation
			HeadingFilter string
		}{
			User:          user,
			Posts:         filteredPosts,
			HeadingFilter: strings.Title(strings.Replace(filterBy, "_", " ", -1)) + " Posts",
		}
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
	r.ParseForm() // 16,1 если авторизованный
	if _, ok := r.Form["filter_by"]; !ok {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w, r)
		return
	}
	value := r.Form["filter_by"]
	if !contains([]string{"oldest", "recent", "most_disliked", "most_liked", "discussions", "questions", "ideas", "articles", "events", "issues", "i_liked", "i_created"}, value[0]) {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w, r)
		return
	}
	filterBy := r.FormValue("filter_by")
	var filteredPosts []model.PostRepresentation
	if filterBy == "i_liked" || filterBy == "i_created" {
		filteredPosts, err = f.serv.PersonalFilter(filterBy, user.ID)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
	} else {
		filteredPosts, err = f.serv.FilterAllPosts((filterBy))
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
	}

	info := struct { // 16,1 эта инфа будет отрисовываться в html
		User          *model.User
		Posts         []model.PostRepresentation
		HeadingFilter string
	}{
		User:          user,
		Posts:         filteredPosts,
		HeadingFilter: strings.Title(strings.Replace(filterBy, "_", " ", -1)) + " Posts",
	}

	t, err := template.New("indexAuthorized.html").Funcs(template.FuncMap{ // 16,1 ваня объяснит
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

func contains(s []string, e string) bool { // 16,1 проверяется есть ли фильтр который пришел с фронта есть в списке "oldest", "recent", "most_disliked", "most_liked", "discussions", "questions", "ideas", "articles", "events", "issues"
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
