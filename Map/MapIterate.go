package main

import "fmt"

func main() {
	my := make(map[string]int)
	my["Hello"] = 1
	my["World"] = 2

	fmt.Println(my)

	for _,value:= range my {
		fmt.Println(value)
	}
}
