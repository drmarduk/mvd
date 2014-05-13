package main

import (
	"fmt"
	"github.com/drmarduk/mvd/NotenSatz"
	"log"
	"net/http"
	 "strconv"
	"strings"
)

func Root(w http.ResponseWriter, r *http.Request) {
	log.Printf("main.Root: request incoming.")
	fmt.Fprint(w, "emtpty request.")
}

func HttpNotenSatz(w http.ResponseWriter, r *http.Request) {
	// /NotenSatz/$action/$params
	path := r.URL.Path
	t := path[len("/NotenSatz/"):]

	t1 := strings.Index(t, "/")
	action := t[:t1]
	params := t[t1+1:]

	log.Printf("Request: %s\nAction: %s\nParams: %s\n", t, action, params)

	if action == "new" {
		r := NotenSatz.New(params, "dir")
		log.Printf("Http New Notensatz: %t\n", r)
	}
	if action == "get" {
		id, err := strconv.Atoi(params)
		if err != nil {
			id = 0
		}
		n, err := NotenSatz.Get(id)
		if err != nil {
			log.Printf("Http Get Notensatz: failed.")
		}
		log.Printf("Http Get NotenSatz: %+v", n)
	}
}
