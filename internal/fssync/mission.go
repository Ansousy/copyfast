package fssync

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

type Mission struct {
	origin string
	target string
	file   fs.DirEntry
}

func (m *Mission) execute() {
	if _, err := os.Stat(m.target + "/" + m.file.Name()); os.IsNotExist(err) {
		fmt.Println("FOLDER:", m.origin+"/"+m.file.Name(), m.target+"/"+m.file.Name())
		os.MkdirAll(m.target+"/"+m.file.Name(), 0777)
	}

	if _, err := os.Stat(m.target + "/" + m.file.Name()); os.IsNotExist(err) {
		fmt.Println("FILE:", m.origin+"/"+m.file.Name(), m.target+"/"+m.file.Name())
		fi, errfi := os.Open(m.origin + "/" + m.file.Name())
		if errfi != nil {
			log.Fatal(errfi)
		}

		//function end up to execute cloture the file automatically
		//whatever it happened it defer the cloture at the end
		defer fi.Close()

		fo, errfo := os.Create(m.target + "/" + m.file.Name())
		if errfo != nil {
			log.Fatal(errfo)
		}

		defer fo.Close()

		if errfi != nil && errfo != nil {
			//only inferior to 1MO
			buffer := make([]byte, 1024)
			for {
				n, err := fi.Read(buffer)
				if err != nil && err != io.EOF {
					log.Fatal(err)
				}

				if n == 0 {
					break
				}

				if _, err := fo.Write(buffer[:n]); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

/*
toto/			1
toto/tata.txt	2
ttoto/titi.txt	1
*/
