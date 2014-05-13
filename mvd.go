package main

import (
	"fmt"
	//"github.com/drmarduk/mvd/NotenSatz"
	"log"
	"net/http"
	"time"
)

var ticker *time.Ticker
var quit chan struct{}

func main() {
	// Setup foo
	//if !NotenSatz.Setup() {
	//	log.Printf("main.main: Failed to create sql tables.")
	//}
	ticker = time.NewTicker(time.Second * 3)
	quit = make(chan struct{})

	go DirListen()

	http.HandleFunc("/", Root)
	//http.HandleFunc("/NotenSatz/", HttpNotenSatz)
	http.ListenAndServe("localhost:8000", nil)
}

func hai() {
	fmt.Println("Hai.")

	for {
		select {
		case <-ticker.C:
			fmt.Println("dörp")
		case <-quit:
			ticker.Stop()
			return
		}
	}

}
