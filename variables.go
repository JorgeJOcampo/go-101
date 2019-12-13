package main

import "fmt"
import "math"

const n = 500

func main(){
	var a, b int = 4, 5
	fmt.Println(a, b)

	//Defaults
	var entero int
	fmt.Println("entero:", entero)
	var string string
	fmt.Println("string:", string)
	fmt.Println("string default:", string == "")
	var boolean bool
	fmt.Println("boolean:", boolean)

	const d = 3e20/n
	fmt.Println("Notacion cientifica", d)
	fmt.Println("En int64", int64(n))

	fmt.Println(math.Sin(n))
	
}