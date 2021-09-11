package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"unsafe"
)

func intit() {
	fmt.Println("Init ... ")
}

func setZeroPassByVal(i int) {
	i = 0
	fmt.Println("Set Zero Pass By Value :", i)
}

func setZeroPassByRef(i *int) {
	*i = 0
	fmt.Println("Set Zero Pass By References :", *i)
}

type developer struct {
	skill		[]string
	exp			int
	salary		int
	location	string
	line 		string
	company 	string
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area 	:", g.area())
	fmt.Println("Perim	:", g.perim())
}

func divide(x ,y int) (int, error) {
	if y <= 0 {
		return 0, errors.New("Can't divide with 0. ")
	}
	return x / y, nil
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
	// For Who learned C you will know it.
	// in Golang you can use Pointer like C

	i := 5 
	fmt.Println(i)
	fmt.Println(&i)
	i = 10
	fmt.Println(i)
	fmt.Println(&i)

	fmt.Println("--------------------------")

	// var <Pointer Name> *<Type of Pointer>
	// <Pointer Name> := <Variable Name>

	// Method 2 Pointer
	 i = 5
	 fmt.Println(" Value of i :", i)
	 fmt.Println(" Address of i :", &i)
	 fmt.Println("--------------------------")
	 
	 p := &i
	 fmt.Println(" Value of P :", p)
	 fmt.Println(" Address of P :", &p)
	 fmt.Println(" Value of *P :", *p)
	 fmt.Println(" Size of P :", unsafe.Sizeof(p), " bytes.")

	 fmt.Println("--------------------------")

	 *p = 10
	 fmt.Println(" Value of *P :", *p)
	 fmt.Println(" Value of i :", i)

	 fmt.Println("--------------------------")
	 // *p is Value of Variable Position Pointer P
	 // &p is Position of P
	
	fmt.Println("--------------------------")
	i = 5
	fmt.Println(" Value of i :", i)
	setZeroPassByVal(i)
	fmt.Println(" Value of i :", i)

	fmt.Println("--------------------------")
	fmt.Println(" Value of i :", i)
	setZeroPassByRef(&i)
	fmt.Println(" Value of i :", i)
	
	fmt.Println("--------------------------")

	
	fmt.Println("--------------------------")
	fmt.Println(" Struct ") // It's From "Structure"

	// Type <Name of Structure> struct {		
	// ...
	// <Name of Variable> <Type of Variable>
	// ...
	// }

	fmt.Println(developer{})
	fmt.Println("--------------------------")
	dev :=  developer {
		skill: 		[]string{"node.js", "Java Springboot"},
		exp:		0,
		location:	"Chaingmai",
		line: 		"n/a",
		company: 	"control alt del",	
	}
	dev.salary = 9000
	fmt.Println(dev)
	fmt.Println("--------------------------")

	devp := &dev
	fmt.Println(devp.company)
	devp.company = "alt  F4"
	fmt.Println(devp.company)
	fmt.Println(dev)
	fmt.Println("--------------------------")

	fmt.Println(" Method Function of Struct ")

	// func (<struct Receiver>) <Name of Function> (Input) <Type of Variable Output> {
	// ...
	// <Code Worker>
	// ...
	// return <result>
	// }

	r := rect{width: 10,height: 5}

	fmt.Println("area :", r.area())
	fmt.Println("perim :", r.perim())
	fmt.Println("--------------------------")

	rp := &r
	fmt.Println("area :", rp.area())
	fmt.Println("perim :", rp.perim())

	fmt.Println("--------------------------")

	// func ( r*rect) area() int {
	// ...
	// <Code Worker>
	// ...
	// return <result>
	// }

	// Interfaces is Structure of Method Inclues. Method Name, Type of Variable and Return

	// Type <Interfaces> interface {
	// ...
	// <Function Name> () <Type of Variable Output>
	// ...
	// }
	fmt.Println("--------------------------")
	fmt.Println(" Interfaces Math Structure Circle ")

	r = rect{width: 3 , height: 4}
	c := circle{radius: 5}

	measure(&r)
	fmt.Println("--------------------------")
	measure(c)

	fmt.Println("--------------------------")
	fmt.Println(" Error Function ")
	// ERROR Return Value another Method
	// errors.New(<Error Message>)
	fmt.Println("--------------------------")

	if r, e := divide(10, 0); e != nil {
		fmt.Println("Error :", e)
	} else {
		fmt.Println("Result :", r)
	}

	fmt.Println("--------------------------")

	}
