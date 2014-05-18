package main

import (
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func DirListen() {
	var errorcountreaddir int = 0
	for {
		log.Printf("New Run.\n")
		files, err := ioutil.ReadDir(InputDir)
		if err != nil {
			log.Printf("main.DirListen: %s\n", err.Error())
			// panic?
			errorcountreaddir++

			if errorcountreaddir > 5 {
				panic(err)
			}
			continue
		}
		log.Printf("[+] %d Items found.", len(files))

		for _, f := range files {
			// process each file
			if !f.IsDir() {
				// nur dateien, keine Ordner
				ProcessFile(f.Name())

			}
		}

		time.Sleep(3 * time.Second)
	}
}

func ProcessFile(file string) {
	log.Println(file)
	name := file[:strings.Index(file, ".")]
	log.Println(name)

}

func MakeDir(name string) bool {
	files, err := ioutil.ReadDir(SaveDir)
	if err != nil {
		log.Printf("main.MakeDir: %s\n", err.Error())
		return false
	}
	for _, n := range files {
		if n.IsDir() {
			if name == n.Name() {
				return true
			}
		}
	}

	return false
}
