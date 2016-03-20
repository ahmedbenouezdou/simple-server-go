package main
import (
	"log"
	"net/http"
)




func main(){


	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/getExemple",GetOnly(getExemple))
	http.HandleFunc("/postExemple",PostOnly(postExemple))

	log.Printf("About to listen on 10443. Go to http://127.0.0.1:9000/ or http://localhost:9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}