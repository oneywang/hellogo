// Q17 project main.go
package main

import (
	"fmt"
)

func main() {
	i := 1
	var p *int = &i
	var j int = *p + 1

	//s := "hi"
	//var q = &s
	//var k = *q + 1 // cannot convert 1 to type string

	//f := false
	//var r = &f
	//var l = *r + 1 // cannot convert 1 to type string

	g := 1.1
	var t = &g
	var m = *t + 1

	//var a []int = make([]int, 10)
	//var u = &a
	//var n = *u + 1	// (mismatched types []int and int

	fmt.Printf("*p++ works for int/float = %v, %v\n!", j, m)
}
