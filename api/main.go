package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
	"encoding/json"
)

//
//STRUCT
//
type Article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article
var (
	flagPort = flag.String("port", "8000", "Port to listen on")
)

// FUNCTION
// FUNCTION
// FUNCTION

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"test",Desc:"TEST",Content:"CISO"},
	}
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(articles)	
		//////fmt.Fprint(w, "GE done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"index hit")
}

func init() {
	log.SetPrefix("LOG: ")
    log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
    log.Println("started")
	flag.Parse()
}
func handleRequests(){

	http.HandleFunc("/",Index)
	http.HandleFunc("/articles",allArticles)
	log.Fatal(http.ListenAndServe(":"+*flagPort,nil))
}

//MAIN
func main(){
	handleRequests()
}