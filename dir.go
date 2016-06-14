package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Eigentlich müste man ein eigenes pkg erstellen, dass einen channel
// annimmt, in den es die gefunden Dateien schieben kann

func DirListen() {
	var errorcountreaddir int = 0
	for {
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
	// file = deutschmeister_regimentsmarsch_0004.pdf
	folder := file[:strings.Index(file, ".")]
	folder = folder[:len(folder)-5]
	// folder = deutschmeister_regimentsmarsch
	
	
	if MakeDir(folder) {
		// try "copy" file to new Dir, via link, does not work on win
		newname := SaveDir + folder + "/" + file
		
		err := os.Link(InputDir+file, newname)
		if err != nil {
			log.Printf("main.ProcessFile: %s\n", err.Error())
			return
		}
		log.Printf("main.ProcessFile: %s moved to %s\n", file, SaveDir+folder)
		err = os.Remove(InputDir + file)
		if err != nil {
			log.Printf("main.ProcessFile: %s\n", err.Error())
			return
		}

	} else {
		log.Printf("main.ProcessFile: failed to create directory for %s\n", folder)
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
	log.Printf("main.MakeDir: %s created.\n", name)
	return true
}
