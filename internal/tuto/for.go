package tuto

import "fmt"

func createFor() {
	//while
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	//for
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//do while
	for {
		fmt.Println("ok")
		if true {
			break
		}
	}

	//names := &[]string{"Florent", "Toto", "Tata"}
	//foreach,_ because index no need
	/*for _, name := range *names {
		fmt.Println(name) //florent, toto, tata
	}*/

	//for index, name := range

}
