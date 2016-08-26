// Learn golang project chapter1.go

package main

import "fmt"

func var_type_declare() {
	var ias []int
	var (
		i uint8
		l bool
		m uint16
	)
	var a int
	var b = 2013
	c := 2014
	const d = 2015
	const (
		e = iota
		f
		g byte = 'p'
		h      = "*shit*"
	)
	i, j, k, _, l, m := 1, 2.08, " golang", false, true, 0xff0
	fmt.Println("Hello World!", len(ias), i, j+5, k, l, m)
	fmt.Println("Get out!", a+2012, b, c, d, e, f, g, h)
}

func string_multiline_str() {
	var my_dreams = "上床 and" +
		" sleep for a while" +
		" and eat a lot"
	var your_dream = `goto heaven
	and fuck virgin`

	// my_dreams[0] = 0x7a // 1. cannot assign 2. it's byte instead of character!
	my_another_dream := []rune(my_dreams)
	my_another_dream[0] = '下'

	fmt.Println(my_dreams)

	fmt.Println(string(my_another_dream))
	// 非
	fmt.Println(string(my_dreams[0]))

	fmt.Println(your_dream)

}

func for_sum(num uint32) uint32 {
	sum := 0
	for i := 0; uint32(i) <= num; i++ {
		sum += i
	}

	sum1, i1 := 0, 0
PLUS:
	sum1 += i1
	i1++
	if uint32(i1) <= num {
		goto PLUS
	}

	if sum1 != sum {
		panic(for_sum)
	}

	return uint32(sum)
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func reverse_string(str string) string {
	arr := []rune(str)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}

func simple_animal(a string) (fly bool, wing bool, feet uint8, swim bool) {
	swim = true
	switch a {
	case "野鸭":
		fly = true
		fallthrough
	case "鸭子":
		wing = true
		feet = 2
	case "青蛙":
		feet = 4
	}
	return
}

func compare(a, b string) int {
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}

	for i := 0; i < len(a); i++ {
		switch {
		case a[i] < b[i]:
			return -1
		case a[i] > b[i]:
			return 1
		}
		continue
	}
	return 0
}

func matrix_multi_vector(a [][]int, b []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a[0]) != len(b) {
		panic(matrix_multi_vector)
	}
	var ret int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			ret += a[i][j] * b[j]
		}
	}
	return ret
}

func array_slice_map(a, b []int) []int {
	if len(a) != len(b) {
		panic(array_slice_map)
	}

	//c需要是个新数组，如果 c := a[:]，或者c := a  则都仅仅是指向a
	var c = make([]int, len(a))
	for k, v := range b {
		c[k] += v
	}
	return c
}
