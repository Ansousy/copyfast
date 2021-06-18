package tuto

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func MakeChannel() {

	messages := make(chan string) //,1) u wait one stuff infinite tunnel
	//Many threads many workers autonomous u can get info from workers from it.

	fmt.Println("1")

	go worker(messages, func() string { return "ping1" })
	go worker(messages, func() string { return "ping2" })
	go worker(messages, func() string { return "ping3" })
	go worker(messages, func() string { return "ping4" })
	fmt.Println("4")

	// when it gets a msg each letter from msg gonna be interpreted in the loop
	for msg := range messages {
		println(string(msg))
	}
	//msg := <-messages //listener
	fmt.Println("5")
	//println(msg)
	runtime.NumCPU()
	//runtime.SetCPUProfileRate(runtime.CPUProfile())

}

//Each worker running between 0 and 3 seconds
func worker(messages chan string, callback func() string) {
	//new independant variable u r receiver
	//<- u get msg and store in msg variables
	//create function that cna be exxecute in asynchronous way
	//go permit to send an function in a asynchronous task
	//go is awaiting an listener and when the listener is online accessible in the api
	fmt.Println("2")
	randomTime := rand.Intn(3) + 1
	for {
		messages <- callback() //emitter
		time.Sleep(time.Duration(randomTime) * time.Second)
	}
	fmt.Println("3")
}

/*
1 Main
\- Thread1 => worker
\- Thread2 => worker
\- Thread3 => worker
*/
/*
//new independant variable u r receiver
	//<- u get msg and store in msg variables
	//create function that cna be exxecute in asynchronous way
	//go permit to send an function in a asynchronous task
	//go is awaiting an listener and when the listener is online accessible in the api
	go func() {
		fmt.Println("2")
		for {
			messages <- "ping" //emitter
			time.Sleep(1 * time.Second)
		}
		fmt.Println("3")
	}()
	fmt.Println("5")
*/
