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
	// curl --request GET http://localhost:2341/files/get?ddh=mary --output saved.zip
	f := r.URL.Query().Get("ddh")
	if f == "" {
		http.Error(w, "parameter missing", http.StatusBadRequest)
		return
	}

	// add folder pre-path and extension, if so
	f = VESSEL_FILES_PATH + f

	// check
	_, e := os.Stat(f)
	if e == nil && strings.HasSuffix(f, ".zip") {
		http.ServeFile(w, r, f)
	} else {
		w.Write([]byte("only files ending in zip"))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/files/get", filesFolderHandler)
	log.Fatal(http.ListenAndServe(":2341", mux))
}
