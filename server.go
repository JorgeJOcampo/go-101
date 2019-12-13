package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/hola", hola)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.Path = %q", r.URL.Path)
}

func hola(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}