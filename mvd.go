package main

import (
	"log"
	"net/http"

	"github.com/drmarduk/mvd/Db"
	"github.com/drmarduk/mvd/NotenSatz"
)

func main() {
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
