package fssync

import (
	"fmt"
	"log"
	"os"
)

func ScanRecursive(origin string, target string) {
	fmt.Println(origin, target)

	if _, err := os.Stat(origin); os.IsNotExist(err) {
		os.Exit(1)
	}

	if _, err := os.Stat(target); os.IsNotExist(err) {
		os.MkdirAll(target, 0777)
	}

	nbMission := scan(origin, target)
	Run()
	//while nbMission is different from current total
	for nbMission != TotalEnCours() {

	}
	fmt.Println("Finish")

}

func scan(origin string, target string) int {
	var nbMission int = 0
	files, err := os.ReadDir(origin)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {

			if file.IsDir() {
				nbMission = scan(origin+"/"+file.Name(), target+"/"+file.Name())
			} else {

				go AddMission(&Mission{
					origin: origin,
					target: target,
					file:   file,
				})
				nbMission++
			}

		}
	}
	return nbMission
}
