package handlers

import (
	"forum/internal/model"
	"html/template"
	"net/http"
	"time"
)

type Authorization interface {
	LoginUser(user *model.User) (uint, bool, int, error)
}

type Session interface {
	SessionCreate(user *model.User) (cookie string, err error)
}

type SingIn struct {
	loginSvc   Authorization
	sessionSvc Session
}

func CreateSignInHandler(loginSvc Authorization, sessionSvc Session) *SingIn {
	return &SingIn{
		loginSvc:   loginSvc,
		sessionSvc: sessionSvc,
	}
}

func (si *SingIn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("authorizedUser").(*model.User)
	if user.ID != 0 { // 20 если авторизован перекидываем на главную страницу
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet { // 20,1 показываем клиенту
		t, err := template.ParseFiles("./templates/signinPage.html")
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
	}
	if r.Method == http.MethodPost { // 20,1 получаем инфу от  клиента если методПост и авторизируем его если все нормально

		err := r.ParseForm()
		if err != nil {
			errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
			return
		}

		userInfo := r.PostForm
		for key := range r.PostForm { // 20,1
			if !contains([]string{"username", "password"}, key) {
				errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
				return
			}
		}
		user := model.User{ // 20,1 получаем юзера
			Username: userInfo["username"][0],
			Password: userInfo["password"][0],
		}
		userId, userLogined, statusCode, err := si.loginSvc.LoginUser(&user) // 20,1  получаем айди юзера
		if err != nil {
			errorPage(err.Error(), statusCode, w, r)
			return
		}
		user.ID = userId
		if userLogined {
			// Create session
			cookie, err := si.sessionSvc.SessionCreate(&user)
			if err != nil {
				errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
				return
			}
			cookieExpiresAt := time.Now().Add(600 * time.Second)
			http.SetCookie(w, &http.Cookie{
				Name:    "Session-token",
				Value:   cookie,
				Expires: cookieExpiresAt,
			})
			// fmt.Println("asd")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		}
		return
	}
	errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w, r)
}
