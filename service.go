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
	io.WriteString(w, "<h1>It works</h1> \n")
}

func getExemple (w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	retrunValue, _ := json.Marshal(Result{"first", "last"})
	io.WriteString(w, string(retrunValue))
}

func postExemple (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.PostFormValue("test"))
	io.WriteString(w, "<h2>post</h2> \n")
}