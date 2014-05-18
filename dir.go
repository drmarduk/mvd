package main

import (
	"io/ioutil"
	"log"
	"os"
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
	// get filename without extension
	// auf_der_vogelwiese_001.pdf
	// deutschmeister_regimentsmarsch_0004.pdf
	name := file[:strings.Index(file, ".")]
	log.Println(name)
	name = name[:len(name)-5]
	log.Println(name)

	if MakeDir(name) {
		// try "copy" file to new Dir
		newname := SaveDir + name + "\\" + file
		log.Printf("Copy to %s\n", newname)
		err := os.Link(InputDir+file, newname)
		if err != nil {
			log.Printf("main.ProcessFile: %s\n", err.Error())
			return
		}

	} else {
		log.Printf("main.ProcessFile: failed to create directory for %s\n", name)
		return
	}

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
	err = os.Mkdir(SaveDir+name, 0777)
	if err != nil {
		log.Printf("main.MakeDir: %s\n", err.Error())
		return false
	}

	return true
}
