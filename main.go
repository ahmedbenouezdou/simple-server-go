package main

import (
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/getExemple",GetOnly(getExemple))
	http.HandleFunc("/postExemple",PostOnly(postExemple))


	log.Fatal(http.ListenAndServe(":8080", nil))
}