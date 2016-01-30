package main

import (
	"net/http"
	"io"
	"encoding/json"
	"log"
)


type Result struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

func HandleIndex (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "It works\n")
}

func getExemple (w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(Result{"tee", "dub"})
	io.WriteString(w, string(result))
}

func postExemple (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.PostForm)
	io.WriteString(w, "post\n")
}