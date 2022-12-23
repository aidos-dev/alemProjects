package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"01.alem.school/git/Optimus/ascii-art-web/ascii"
)

func main() {
	HandleRequest()
}

func HandleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", homePage)
	log.Println("Listening...")
	http.ListenAndServe(":4000", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		// errorHandler(w, r, http.StatusNotFound)
		log.Print("\nError: Not found 404\n")
		return
	}
	tmpl, err := template.ParseFiles(
		"templates/header.html",
		"templates/footer.html",
		"templates/home.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// bannerInp is a variable which stores a selected radio button value of 3 fonts:
	// standard, shadow, thinkertoy
	// bannerInp := "standard"
	bannerInp := r.FormValue("banner")

	// inputArg is a variable which stores the input value from the user as a string
	// FormValue func parses input text from the home page text area and stores it to inputArg variable
	// inputArg := ""
	inputArg := r.FormValue("inputString")

	tmpl.ExecuteTemplate(w, "home", nil)

	Output, err := ascii.AsciiConv(inputArg, bannerInp)
	if err != nil {
		log.Printf("Bad request error (400), %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		// http.Error(w, err.Error(), 400)
		return
	}

	// all "\n" symbols replaced with "<br>" because that is how web pages recognize new lines
	Output = strings.ReplaceAll(Output, "\n", "<br>")

	// Output variable is wrapped with <pre></pre> tags to keep the original format.
	// without this tag a web page ignores some white spaces and it makes the output incorrect
	Output = "<pre>" + Output + "</pre>"

	// Final web output
	fmt.Fprintf(w, Output)

	// reseting variables to default values after each convertion
	// in order not to repeat old values to the web page and terminal after
	// the page refresh
	inputArg = ""
	err = nil

	// Testing section:
	//
	//
	// test terminal output
	// fmt.Println(Output)
	// // test terminal output
	// switch bannerInp {
	// case "standard":
	// 	fmt.Println("standard")
	// case "shadow":
	// 	fmt.Println("shadow")
	// case "thinkertoy":
	// 	fmt.Println("thinkertoy")
	// }

	// // test terminal output
	// fmt.Println(inputArg + " - test")

	// test web output
	// fmt.Fprintf(w, "<pre>Hello <br> new  \n wo            rld \n\n</pre>")
}

// func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
// 	w.WriteHeader(status)
// 	if status == http.StatusNotFound {
// 		fmt.Fprint(w, "\n\nOhh... no... it is - 404 error :(")
// 	}
// }
