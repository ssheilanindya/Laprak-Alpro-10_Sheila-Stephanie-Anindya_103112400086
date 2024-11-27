package main

import "fmt"

func main() {

	var a int
	var prima bool = (true)

	fmt.Print("bilangan : ")
	fmt.Scan(&a)

	fmt.Print("faktor : ")

	for i := 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Print(i, " ")
		}
	}
	for i := 2; i < a; i++ {
		if a%i == 0 {
			prima = false
		}
	}
	fmt.Println("")
	fmt.Println("Prima : ", prima)
}
