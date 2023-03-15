package handlers

import (
	"forum/internal/model"
	"net/http"
	"time"
)

type Leaving interface {
	DeleteSession(userId uint) error
	UserBySession(cookie string) (*model.User, error)
}

type Logout struct {
	service Leaving
}

func CreateLogoutHandler(service Leaving) *Logout {
	return &Logout{
		service: service,
	}
}

func (l Logout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorPage(http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, w, r)
		return
	}

	if r.URL.Path != "/logout" {
		errorPage(http.StatusText(http.StatusNotFound), http.StatusNotFound, w, r)
		return
	}

	user := r.Context().Value("authorizedUser").(*model.User)
	if user.ID == 0 {
		errorPage(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, w, r)
		return
	}

	// fmt.Println("LOGOUT")
	err := l.service.DeleteSession(user.ID)
	if err != nil {
		//	fmt.Println("2")
		errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Session-token",
		Value:   "",
		Expires: time.Now(),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
