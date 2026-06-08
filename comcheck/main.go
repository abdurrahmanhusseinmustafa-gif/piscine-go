package main

import (
	"fmt"
	"os"
)

func main() {
	for _, r := range os.Args[1:] {
		if r == "01" || r == "galaxy" || r == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
