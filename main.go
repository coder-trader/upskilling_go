package main

import "fmt"

func main() {
	for i := range 10 {
		defer fmt.Println(i)
	}
}
