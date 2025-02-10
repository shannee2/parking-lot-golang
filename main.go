package main

import "fmt"

type A interface {
}

func main() {
	s, b := A.(string)
	fmt.Println(s, b)
}
