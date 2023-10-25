package main

import "fmt"

type Vertex struct{
	number int
}

func (v *Vertex) print(){
	fmt.Println(v.number)
	v.number=10
}

func main() {
	fmt.Println(a...,"Hi")
}
