package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"01.alem.school/git/Optimus/groupie-tracker-visualizations/internals/groupie"
)

// The ParsedJson variable is created as global in order to parse JSON
// only once and then have imidiate access to JSON data from any handler of function
var ParsedJson groupie.GeneralApiStruct

// The Id variable created as global in order to keep it alive after parsing
var Id int

func main() {
	// groupie.TestParser()

	ApiParser, err := groupie.ParseJson()
	if err != nil {
		log.Println(err)
		return
	}

	ParsedJson = ApiParser

	HandleRequest()
}

func HandleRequest() {
	// Router
	router := http.NewServeMux()
	// files server
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// home page
	router.HandleFunc("/", HomePage)
	// groups details page
	router.HandleFunc("/details/", DetailsPage)
	// Server
	server := http.Server{
		Addr:         ":4000",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	// Run Server
	log.Println("Listening on http://localhost:4000/\n")
	err := server.ListenAndServe()
	if err != nil {
		log.Print(err)
		return
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/favicon.ico" {
		ErrorHandler(w, http.StatusNotFound, "404 Page Not Found")
		return
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Internal server error. Status: 500", http.StatusInternalServerError)
		log.Printf("home handler template parsing error: %v", err)
		return
	}

	if r.Method != "GET" {
		fmt.Println(r.Method)
		ErrorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed. Status: 405")
		return
	}

	Data := ParsedJson.GroupsAll

	tmpl.Execute(w, Data)
}

func DetailsPage(w http.ResponseWriter, r *http.Request) {
	parseId, _ := strconv.Atoi(r.URL.Query().Get("id"))

	// This condition is required since Query parses id twice and on the second run makes it 0.
	// It leads to panic since there is no such id in json structure
	if parseId != 0 {
		Id = parseId
	}

	linkCheck := "/details/?id=" + strconv.Itoa(Id)

	if (r.URL.Path != "/details/" && r.URL.Path != "/details/static/favicon/favicon.ico/") || r.RequestURI != linkCheck || Id > len(ParsedJson.GroupsAll) || Id < 0 {
		ErrorHandler(w, http.StatusNotFound, "404 Page Not Found")
		return
	}

	tmpl, err := template.ParseFiles("templates/details.html")
	if err != nil {
		http.Error(w, "Internal server error. Status: 500", http.StatusInternalServerError)
		log.Printf("details handler template parsing error: %v", err)
		return
	}

	if r.Method != "GET" {
		fmt.Println(r.Method)
		ErrorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed. Status: 405")
		return
	}

	Data := groupie.GroupDetails{
		Group:     ParsedJson.GroupsAll[Id-1],
		Locations: ParsedJson.LocationsAll.LocsGeneral[Id-1],
		Dates:     ParsedJson.DatesAll.Dates[Id-1],
		Relations: ParsedJson.RelationsAll.IndexRels[Id-1],
	}

	tmpl.Execute(w, Data)
}

func ErrorHandler(w http.ResponseWriter, status int, errMessage string) {
	tmpl, err := template.ParseFiles(
		"templates/errorPages/errorPage.html",
	)
	if err != nil {
		http.Error(w, "internal server error. Status: 500", http.StatusInternalServerError)
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
