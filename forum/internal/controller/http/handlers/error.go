package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func errorPage(errorType string, code int, w http.ResponseWriter, r *http.Request) { // 15 функция вызывается при ошиках
	w.WriteHeader(code)
	fmt.Printf("%s %s [%s]\t%s%s - %d - %s\n", time.Now().Format("2006/01/02 15:04:05"), r.Proto, r.Method, r.Host, r.RequestURI, code, http.StatusText(code)) // 15 если есть ошибка, принтится в консоль, /r  указатель мыши ставит в начало и переписывает все что было выше заново
	fmt.Println(errorType)
	t, err := template.ParseFiles("./templates/error.html") // 15,1 страница ошибки показывается
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusText(http.StatusInternalServerError))
	}
	data := struct {
		Err  string
		Code int
	}{
		Err:  errorType,
		Code: code,
	}

	err = t.Execute(w, data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusText(http.StatusInternalServerError))
	}
}
