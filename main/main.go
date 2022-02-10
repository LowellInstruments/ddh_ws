package main

// the / path is a subtree path (trailing /)
// the about path is a fixed path

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// otherwise, even bad urls are served with their most-common
	// route existing in the server, even root one
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("<h1>Welcome to my web server!</h1>"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>about page!</h1>"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// ex: http://localhost:8080/user?id=123&sex='m'
	id := r.URL.Query().Get("id")
	sex := r.URL.Query().Get("sex")
	if id == "" {
		http.Error(w, "The id query parameter is missing", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "<h1>id is: %s</h1>", id)
	fmt.Fprintf(w, "<h1>sex is: %s</h1>", sex)
}

func filesFolderHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/files/get?ddh=mary
	ddh := r.URL.Query().Get("ddh")
	if ddh == "" {
		http.Error(w, "ddh parameter is missing", http.StatusBadRequest)
		return
	}
	//fmt.Fprintf(w, "<h1>ddh is: %s</h1>", ddh)

	http.ServeFile(w, r, "goku.png")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/user", userHandler)
	mux.HandleFunc("/files/get", filesFolderHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

// curl examples
// curl --request GET http://localhost:8080/user?id=123&sex='m'
// curl --request GET http://localhost:8080/files/get?ddh=mary --output saved.zip
