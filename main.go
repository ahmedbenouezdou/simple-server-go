package main

import (
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/getExemple",GetOnly(getExemple))
	http.HandleFunc("/postExemple",PostOnly(postExemple))

	log.Printf("About to listen on 10443. Go to https://127.0.0.1:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}