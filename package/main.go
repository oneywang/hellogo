package main

import (
	"fmt"
	"helloworld/even"
)

func calcu2(a, b, c int) (int, string) {
	ret, s := 0, ""
	if a+b == c {
		s = fmt.Sprintf("%v + %v", a, b)
		ret += 1
	}
	if a*b == c {
		s = fmt.Sprintf("%v * %v", a, b)
		ret += 1
	}
	if a-b == c {
		s = fmt.Sprintf("%v - %v", a, b)
		ret += 1
	}
	if a%b == 0 && a/b == c {
		s = fmt.Sprintf("%v / %v", a, b)
		ret += 1
	}
	return ret, s
}

func select2(p []int, result int) int {
	var plen = len(p)
	if plen < 2 {
		panic(select2)
	}

	count := 0
	for i := 0; i < plen; i++ {
		for j := 0; j < plen; j++ {
			if has, s := calcu2(p[i], p[j], result); has > 0 {
				count += has
				fmt.Printf("%s = %v\n", s, result)
			}
		}
	}
	return count
}

func calcu3(a, b, c, d int, debug bool) (int, string) {
	has, s := 0, ""
	if d%a == 0 {
		if has, s = calcu2(b, c, d/a); has > 0 {
			s = fmt.Sprintf("%v * (%v)", a, s)
			ret += 1
			if debug {
				fmt.Printf("%v = %v\n", s, d)
			}
		}
	}
	if has, s = calcu2(b, c, d*a); has > 0 {
		s = fmt.Sprintf("(%v) / %v", s, a)
		if debug {
			fmt.Printf("%v = %v\n", s, d)
		}
	}
	if has, s = calcu2(b, c, d-a); has > 0 {
		s = fmt.Sprintf("%v + %v", a, s)
		if debug {
			fmt.Printf("%v = %v\n", s, d)
		}
	}
	if has, s = calcu2(b, c, d+a); has > 0 {
		s = fmt.Sprintf("%v - %v", s, a)
		if debug {
			fmt.Printf("%v = %v\n", s, d)
		}
	}

	return has, s
}

func select3(p []int, result int) int {
	var plen = len(p)
	if plen < 3 {
		panic(select3)
	}

	count := 0
	for i := 0; i < plen; i++ {
		for j := 0; j < plen; j++ {
			for k := 0; k < plen; k++ {
				if has, _ := calcu3(p[i], p[j], p[k], result); has > 0 {
					count += has
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println(Even(101))
	fmt.Println(even.Even(101))

	var p = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 25, 50, 75, 100}
	result := 22
	fmt.Printf("Total %v! \n", select2(p, result))
	fmt.Printf("Total %v! \n", select3(p, result))

}
