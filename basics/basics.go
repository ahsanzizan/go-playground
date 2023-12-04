package main

import "fmt"

var a int // Single Variable

var (
	b bool
	c float32
	d string
	e uint32
)

func main() {
	a = 69
	b, c = true, 32.9
	d = "Hi Mom"
	fmt.Println(a, b, c, d, e)
}
