package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"01.alem.school/git/Optimus/ascii-art-web-stylize/ascii"
)

func main() {
	HandleRequest()
}

func HandleRequest() {
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// http.HandleFunc("/", homePage)

	// log.Println("Listening on http://localhost:4000/\n")

	// err := http.ListenAndServe(":4000", nil)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }

	// Router
	router := http.NewServeMux()

	// files server
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// home page
	router.HandleFunc("/", homePage)

	// Server
	server := http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	// Run Server
	log.Println("Listening on http://localhost:4000/\n")
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
	if r.URL.Path != "/" && r.URL.Path != "/favicon.ico" {
		// http.NotFound(w, r)
		errorHandler(w, http.StatusNotFound, "404 Not Found")
		// log.Print("\nError: Not found 404\n")
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/header.html",
		"templates/footer.html",
		"templates/home.html",
		// "templates/errorPages/errorPage.html",
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

	// bannerInp is a variable which stores a selected radio button value of 3 fonts:
	// standard, shadow, thinkertoy
	// bannerInp := "standard"
	bannerInp := r.FormValue("banner")

	// inputArg is a variable which stores the input value from the user as a string
	// FormValue func parses input text from the home page text area and stores it to inputArg variable
	// inputArg := ""
	inputArg := r.FormValue("inputString")
	// bodyText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		return
	}

	// Testing terminal print outs
	// fmt.Println(r.Method)
	// fmt.Println(inputArg)
	// fmt.Println(len(inputArg))

	if r.Method != "POST" && r.Method != "GET" {
		fmt.Println(r.Method, inputArg)
		errorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed. Status: 405")
		return
	}

	// http.Redirect(w, r, "/ascii/", http.StatusFound)

	Output, err, errHttpCode := ascii.AsciiConv(inputArg, bannerInp)
	if err != nil {
		// log.Printf("Bad request error (400), %v", err)
		// fmt.Fprintf(w, "Error: Bad request. Status: 400")
		if errHttpCode == 400 {
			errorHandler(w, http.StatusBadRequest, "400 Bad Request")
		}
		if errHttpCode == 500 {
			errorHandler(w, http.StatusInternalServerError, "Internal server error. Status: 500")
		}
		// http.Error(w, err.Error(), http.StatusBadRequest)
		// log.Print("\nError: Bad request 400\n")
		// http.Redirect(w, r, "/errorPage/", http.StatusSeeOther)

		return
		// http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		// return
	}
	// all "\n" symbols replaced with "<br>" because that is how web pages recognize new lines
	// Output = strings.ReplaceAll(Output, "\n", "<br>")

	// Output variable is wrapped with <pre></pre> tags to keep the original format.
	// without this tag a web page ignores some white spaces and it makes the output incorrect
	// Output = "<pre>" + Output + "</pre>"

	tmpl.ExecuteTemplate(w, "home", Output)

	// Final web output
	// condition "if len(Output)>11" is required to perform clear printing of data
	// in a terminal window without unnecessary empty repeatitions
	// if len(Output) > 11 {
	// 	fmt.Fprintf(w, Output)
	// 	// fmt.Printf("banner: [%v]\n", bannerInp)
	// 	// fmt.Printf("input: [%v]\n", inputArg)
	// 	// fmt.Printf("lenth: [%v]\n", len(Output))
	// 	// fmt.Printf("netLenth: [%v]\n", len(Output)-11)
	// }

	// r.ParseForm()
	// text, ok1 := r.Form["inputString"]
	// bannerInp, ok2 := r.Form["banner"]

	// fmt.Printf("input: %v\n\n", text)
	// fmt.Printf("banner: %v\n\n", bannerInp)
	// fmt.Printf("ok1: %v\n\n", ok1)
	// fmt.Printf("ok2: %v\n\n", ok2)

	// if !ok1 || !ok2 {
	// 	errorHandler(w, http.StatusBadRequest, "400 Bad Request")
	// 	return
	// }

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

func errorHandler(w http.ResponseWriter, status int, errMessage string) {
	tmpl, err := template.ParseFiles(
		"templates/errorPages/errorPage.html",
	)
	if err != nil {
		http.Error(w, "Internal server error. Status: 500", http.StatusInternalServerError)
		log.Printf("Error: handler template parsing error: %v", err)
		return
	}

	w.WriteHeader(status)
	err = tmpl.Execute(w, errMessage)

	if err != nil {
		log.Printf("Error: handler template execution error: %s", err.Error())
	}
	log.Print(errMessage)
}
