package main

import (
	"net/http"
	"io"
	"log"
	"io/ioutil"
	"fmt"
	"os"
)




//index
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>It's works</h1> \n")
}

/**
function get
 */
func getExemple(w http.ResponseWriter, r *http.Request) {

	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(string(file)))
}

/**
function post
 */
func postExemple(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.PostFormValue("test"))
	io.WriteString(w, "<h2>post</h2> \n")
}