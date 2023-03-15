package handlers

import (
	"context"
	"fmt"
	"forum/internal/model"
	"net/http"
	"time"
)

type Auth interface {
	SessionCheck(cookie string) (bool, error)
	UserBySession(cookie string) (*model.User, error)
	DeleteSession(userId uint) error
}

type Middleware struct {
	service Auth
}

func CreateMiddleware(service Auth) *Middleware {
	return &Middleware{
		service: service,
	}
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Session-token") // 11,1 получение куки из реквеста
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "authorizedUser", &model.User{})))
			return
		}

		user, err := m.service.UserBySession(cookie.Value) // 11,1 нахождение юзера по значению куки (пока пустой юзер)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "authorizedUser", &model.User{}))) //(в реквесте пустой юзер)
			return
		}

		if cookie.Expires.After(time.Now()) { // 11.1 проверка на время жизни куки, если оно прошло, то удаляется сессия
			err = m.service.DeleteSession(user.ID)
			if err != nil {
				errorPage(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, w, r)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "authorizedUser", &model.User{})))
			return
		}

		cookieExpiresAt := time.Now().Add(600 * time.Second) // 11,1 здесь обновляется время жизни куки 10 минут пока юзер запросы делает

		http.SetCookie(w, &http.Cookie{ // 11,1 здесь сохраняет куки у клиента
			Name:    "Session-token",
			Value:   cookie.Value,
			Expires: cookieExpiresAt,
		})

		ctx := context.WithValue(r.Context(), "authorizedUser", user) // 11,1 на основе реквест контекста создается новый контекст в который записываается юзер под ключом "authorizedUser"

		next.ServeHTTP(w, r.WithContext(ctx)) // 11,1 в реквест записывается этот контекст уже с юзером
	})
}

func (m *Middleware) LoggerMiddlerware(next http.Handler) http.Handler { // 11,1 стандартное логгирование показывает в терминале какой метод какой хост какой эндпоинт
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s [%s]\t%s%s - 200 - OK\n", time.Now().Format("2006/01/02 15:04:05"), r.Proto, r.Method, r.Host, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// FIXME:
