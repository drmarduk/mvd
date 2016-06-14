package main

import (
	"github.com/drmarduk/mvd/NotenSatz"
	"github.com/drmarduk/mvd/Db"
	"log"
	"net/http"
)

func main() {
	log.Printf("Start.")

	log.Printf("Setup Database.")
	Db.Hi()
	// dir foo
	if RunDir {
		log.Printf("Start DirListen().\n")
		go DirListen()
	}

	if RunHttp {
		log.Printf("Start HttpListen().\n")
		// Setup foo
		if !NotenSatz.Setup() {
			log.Printf("main.main: Failed to create sql tables.")
		}
		// http handler foo
		http.HandleFunc("/", Root)
		http.HandleFunc("/NotenSatz/", HttpNotenSatz)
		http.ListenAndServe("localhost:8000", nil)
	}
	
}
