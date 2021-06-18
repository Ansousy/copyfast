package fssync

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

func CopyRecursive(origin string, target string) {
	fmt.Println(origin, target)

	if _, err := os.Stat(origin); os.IsNotExist(err) {
		os.Exit(1)
	}

	if _, err := os.Stat(target); os.IsNotExist(err) {
		os.MkdirAll(target, 0777)
	}

	copy(origin, target)
}

func copy(origin string, target string) {

	files, err := os.ReadDir(origin)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {

			if file.IsDir() {
				copyFolder(origin, target, file)
			} else {
				copyFile(origin, target, file)
			}

		}
	}
}

func copyFolder(origin string, target string, file fs.DirEntry) {
	if _, err := os.Stat(target + "/" + file.Name()); os.IsNotExist(err) {
		fmt.Println("FOLDER:", origin+"/"+file.Name(), target+"/"+file.Name())
		os.MkdirAll(target+"/"+file.Name(), 0777)
	}
	copy(origin+"/"+file.Name(), target+"/"+file.Name())
}

func copyFile(origin string, target string, file fs.DirEntry) {
	if _, err := os.Stat(target + "/" + file.Name()); os.IsNotExist(err) {
		fmt.Println("FILE:", origin+"/"+file.Name(), target+"/"+file.Name())
		fi, errfi := os.Open(origin + "/" + file.Name())
		if errfi != nil {
			log.Fatal(errfi)
		}

		//function end up to execute cloture the file automatically
		//whatever it happened it defer the cloture at the end
		defer fi.Close()

		fo, errfo := os.Create(target + "/" + file.Name())
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
