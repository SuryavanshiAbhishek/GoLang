// Wait Gorups in Go

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func GetData(s string) {
	defer wg.Done()
	res, err := http.Get(s)

	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println("Success", res.StatusCode)
}

func main() {


	links := []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for i := 0; i < len(links); i++ {
		go GetData(links[i])
		wg.Add(1)
	}

	wg.Wait()
}
