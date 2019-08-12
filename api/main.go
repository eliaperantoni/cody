package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"./conn"
	"github.com/blevesearch/bleve"
	"github.com/gorilla/mux"
)

var idx bleve.Index

var flagPort = flag.String("port", "8000", "Port to listen on")

func init() {
	log.SetPrefix("init: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("init")
	flag.Parse()
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
	log.Println("index")
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":"+*flagPort, myRouter))
}

//MAIN
const (
	testIdx = "test.index"
)

func main() {
	idx, err := conn.BlaveIndex(testIdx)
	if err != nil {
		log.Println("Errore")
		log.Println(idx)
	}
	r, err := conn.FindByText(idx, "ciao")
	if err != nil {
		log.Println("Errore")
	} else {
		//ELEMENTI TROVATI
		for i := 0; i < int(r.Total); i++ {
			log.Println(r.Hits[i].ID)

		}
	}

	//d := models.InitCard("1.md")
	//err = d.Index(idx)
	//log.Println(err == nil)
	//handleRequests()

}
