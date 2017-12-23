package main

import "fmt"

	func main() {

		fmt.Println("ASCII Table")
		fmt.Println("-----------")
		var i int = 32
		for i <= 126 {
			fmt.Printf("%d - %b - %x -- %c\n",i,i,i,i)
			i = i + 1
		}


	}
