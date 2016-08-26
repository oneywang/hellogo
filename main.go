// Learn golang project main.go
package main

import "fmt"

func chapter1() {
	fmt.Println("Chapter1 Demo:")
	var_type_declare()
	string_multiline_str()
	fmt.Println("1...100 total =", for_sum(100))
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(reverse(arr))

	fmt.Println(simple_animal("鸭子"))
	fmt.Println(simple_animal("野鸭"))

	s1, s2, s3 := "you are light", "you are a", "I am golang geek!"
	s4 := s3
	fmt.Println(compare(s1, s2), compare(s1, s3), compare(s3, s4))

	a1, a2 := []int{1, 2}, []int{3, 4}
	fmt.Println(array_slice_map(a1, a2))

	a3 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(a3, "*", a1, "=", matrix_multi_vector(a3, a1))

	fmt.Println("Chapter1 Homework: ")
	fmt.Println(fab(99))
	/*for i := 1; i <= 60; i++ {
		fmt.Println(fab(uint32(i)))
	}*/

	s := "1234五六七8九十"
	fmt.Println(s)
	fmt.Println(reverse_string(s))

	for i, s := 0, ""; i < 10; i++ {
		s += "A"
		fmt.Println(s)
	}
}

func rec(i int) {
	//defer fmt.Println(i)
	if i == 10 {
		return
	}

	s := ""
	for j := 0; j < i; j++ {
		s += "牛"
	}
	fmt.Println(s)
	defer fmt.Println(s)

	rec(i + 1)
}

func printit(x int) {
	fmt.Printf("%v\n", x)
}

func chapter2() {
	fmt.Println("Chapter2 Demo:")
	fmt.Println(identity(100))

	rec(1)

	fmt.Println(myfunc_args(2, 1, 2))

	callback_p(999, printit)
	callback_p(999, func(in int) {
		fmt.Println(in)
	})

	fmt.Println("Chapter2 Homework:")
	var arr_fl64 = []float64{1.3, 0.21, 3}
	fmt.Println(average(arr_fl64))

	fmt.Println(ordered_pair(7, 2))
	fmt.Println(ordered_pair(2, 7))

	fmt.Println(
		map_it(
			func(x int) int {
				return x + 100
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7}))

	fmt.Println(bubble_sort([]int{9, 8, 3, 2, 5, 1}))

	fmt.Println(plusTwo()(999))
	fmt.Println(plusX(1)(999))
}

func test() {
	var p = []int{1, 2, 3}
	//if p[3] {
	fmt.Println("add new")
	p[3] = 4
	//}
	fmt.Println(p)
}
func main() {
	test_interface1()
	return
	chapter2()

}
