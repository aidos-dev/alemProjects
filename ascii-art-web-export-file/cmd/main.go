package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"01.alem.school/git/Optimus/ascii-art-web-export-file/ascii"
)

/*
this variable (outputGlobal) is created as global in order to hold data (string) from user ascii-art output
and write it to a .txt file when user clicks the "save" button
*/
var outputGlobal string

func main() {
	HandleRequest()
}

func HandleRequest() {
	// Router
	router := http.NewServeMux()

	// files server
	router.Handle("/webFiles/static/", http.StripPrefix("/webFiles/static/", http.FileServer(http.Dir("./webFiles/static/"))))

	// home page
	router.HandleFunc("/", homePage)

	// "save" button handler
	router.HandleFunc("/save", save)

	// Server
	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Run Server
	log.Println("Listening on http://localhost:8080/\n")
	err := server.ListenAndServe()
	if err != nil {
		log.Print(err)
		return
	}
}

// global variables for text input and banner
// they are created as global so different handlers can have access to them
// var inputArg, bannerInp string

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/webFiles/static/favicon.ico" {
		// http.NotFound(w, r)
		errorHandler(w, http.StatusNotFound, "404 Not Found")
		// log.Print("\nError: Not found 404\n")
		return
	}

	tmpl, err := template.ParseFiles(

		"webFiles/templates/header.html",
		"webFiles/templates/footer.html",
		"webFiles/templates/home.html",
	)
	if err != nil {
		http.Error(w, "Internal server error. Status: 500", http.StatusInternalServerError)
		log.Printf("home handler template parsing error: %v", err)
		return
	}

	// text, ok :=

	// u, err := url.ParseRequestURI("http://localhost:4000/")
	// fmt.Println(u)
	// fmt.Println(u.Scheme)
	// fmt.Println(u.Host)

	/*
		bannerInp is a variable which stores a selected radio button value of 3 fonts:
		standard, shadow, thinkertoy
		bannerInp := "standard"
	*/
	bannerInp := r.FormValue("banner")

	/*
		inputArg is a variable which stores the input value from the user as a string
		FormValue func parses input text from the home page text area and stores it to inputArg variable
		inputArg := ""
	*/
	inputArg := r.FormValue("inputString")

	if err != nil {
		log.Print(err)
		return
	}

	if r.Method != "POST" && r.Method != "GET" {
		fmt.Println(r.Method, inputArg)
		errorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed. Status: 405")
		return
	}

	Output, err, errHttpCode := ascii.AsciiConv(inputArg, bannerInp)
	if err != nil {

		if errHttpCode == 400 {
			errorHandler(w, http.StatusBadRequest, "400 Bad Request")
		}
		if errHttpCode == 500 {
			errorHandler(w, http.StatusInternalServerError, "Internal server error. Status: 500")
		}

		return

	}

	outputGlobal = Output

	tmpl.ExecuteTemplate(w, "home", Output)
}

func save(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println(r.Method)
		errorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed. Status: 405")
		return
	}
	var write []byte
	write = append(write, []byte(outputGlobal)...)

	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(outputGlobal)))
	err := ioutil.WriteFile("webFiles/static/files/ascii-art.txt", write, 0o644)
	if err != nil {
		fmt.Println(err)
	}

	http.ServeFile(w, r, "webFiles/static/files/ascii-art.txt")
}

func errorHandler(w http.ResponseWriter, status int, errMessage string) {
	tmpl, err := template.ParseFiles(
		"webFiles/templates/errorPages/errorPage.html",
	)
	if err != nil {
		http.Error(w, "Internal server error. Status: 500", http.StatusInternalServerError)
		log.Printf("error: handler template parsing error: %v", err)
		return
	}

	w.WriteHeader(status)
	err = tmpl.Execute(w, errMessage)

	if err != nil {
		log.Printf("error: handler template execution error: %s", err.Error())
	}
	log.Print(errMessage)
}
