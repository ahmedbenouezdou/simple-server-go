package main

import (
    "net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

func GetOnly(h handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            h(w, r)
            return
        }
        http.Error(w, "get serve only", http.StatusMethodNotAllowed)
    }
}

func PostOnly(h handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            h(w, r)
            return
        }
        http.Error(w, "post serve only", http.StatusMethodNotAllowed)
    }
}

func DeleteOnly(h handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "DELETE" {
            h(w, r)
            return
        }
        http.Error(w, "Delete serve only", http.StatusMethodNotAllowed)
    }
}

func PutOnly(h handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "PUT" {
            h(w, r)
            return
        }
        http.Error(w, "put serve only", http.StatusMethodNotAllowed)
    }
}

