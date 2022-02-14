package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var VESSEL_FILES_PATH = "files/"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// only exactly '/' or this would act as fall-back
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files, err := ioutil.ReadDir(VESSEL_FILES_PATH)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		f := file.Name()
		if file.IsDir() || !strings.HasSuffix(f, ".zip") {
			continue
		}

		s := file.Name() + "$"
		w.Write([]byte(s))
	}
}

func filesFolderHandler(w http.ResponseWriter, r *http.Request) {
	// curl --request GET http://localhost:8080/files/get?ddh=mary --output saved.zip
	f := r.URL.Query().Get("ddh")
	if f == "" {
		http.Error(w, "parameter missing", http.StatusBadRequest)
		return
	}

	// add folder pre-path
	f = VESSEL_FILES_PATH + f
	if !strings.HasSuffix(f, ".zip") {
		f += ".zip"
	}

	// check
	_, e := os.Stat(f)
	if e == nil {
		http.ServeFile(w, r, f)
	} else {
		w.Write([]byte("<h1>nope file</h1>"))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/files/get", filesFolderHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
