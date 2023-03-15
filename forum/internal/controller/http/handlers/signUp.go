package handlers

import (
	"forum/internal/model"
	"html/template"
	"net/http"
)

type Registration interface {
	RegisterUser(user *model.User) (int, error)
}

type SignUp struct {
	service Registration
}

func CreateSignUpHandler(service Registration) *SignUp {
	return &SignUp{
		service: service,
	}
}

func (su *SignUp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("authorizedUser").(*model.User)
	if user.ID != 0 { // 21 еслти авторизоавн чтобы не видел страницу авторизации
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("./templates/signupPage.html")
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
			return
		}
		return
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		userInfo := r.PostForm
		for key := range r.PostForm {
			if !contains([]string{"email", "username", "password"}, key) {
				errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
				return
			}
		}
		user := model.User{
			Email:    userInfo["email"][0],
			Username: userInfo["username"][0],
			Password: userInfo["password"][0],
		}
		statusCode, err := su.service.RegisterUser(&user) // вызывваем метод который регистрирует на сервисе, а с серивса он вызывает репозиторий
		if err != nil {
			errorPage(err.Error(), statusCode, w, r)
			return
		}
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}
	errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w, r)
}
