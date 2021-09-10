package main

import (
	"fmt"
	"os"
)

func intit() {
	fmt.Println("Init ... ")
}

func main() {
	fmt.Println("Day#6 Start")
	fmt.Println("--------------------------")

	// Recover() Function is manage runtime error from Panic()
	fmt.Println(" Recover Function ")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(" Recovered : The System cannot find the path specified. ")
			fmt.Println("--------------------------")
		}
	}()

	// Defer is Last in First or Last to First

	defer fmt.Println("After end of the main() func.")
	defer fmt.Println("Time for Tubby Bye-Bye!")
	defer fmt.Println("--------------------------")

	fmt.Println(" Hello Gopher. ")
	fmt.Println("--------------------------")

	fmt.Println(" Panic Function ")

	// Panic is STOP Working! Function
	// When you was wrong in your code
	// but if you don't want it run on your Product.

	path := "/Users/Harit/golang-training2/Password.txt"
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	fmt.Println(" ==> Done creating file. ", path)

	fmt.Println("--------------------------")

	fmt.Println(" Pointer Function ")

	// Pointer 
	

}
