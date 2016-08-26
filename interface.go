package main

import "fmt"

type Interface interface {
	String() string
}

type Implementation int

func (v Implementation) String() string {
	return fmt.Sprintf("Hello %d", v)
}

func test_interface1() {
	var i Interface
	impl := Implementation(42)
	i = impl
	fmt.Println(i.String())
	impl = Implementation(91)
	fmt.Println(i.String())
}
