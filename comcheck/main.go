package main

import (
	"fmt"
	"os"
)

func main() {
	size := len(os.Args)
	if size < 2 {
		return
	}

	counter := 0
	for counter < size && os.Args[counter] != "01" && os.Args[counter] != "galaxy" && os.Args[counter] != "galaxy 01" {
		counter++
	}

	if counter < size {
		fmt.Println("Alert!!!")
	}
}
