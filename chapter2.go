package main

// 输入参数不能与输出参数重名！Error: duplicate argument "in"
// func identity(in int) (in int) {
func identity(in int) (out, out1 int) {
	out = in
	out1 = in + 1
	return
	//return out, out1 //多返回值就这么简单地用,分别列出，注意次序即可！
}

func myfunc_args(arg ...int) []int {
	//var ret = arg[:]
	return arg
}

func callback_p(y int, f func(int)) {
	f(y)
	return
}

// Homework

func average(in []float64) float64 {
	count := len(in)
	if count == 0 {
		return float64(0)
	}

	var sum float64
	for i := 0; i < count; i++ {
		sum += in[i]
	}
	return sum / float64(count)
}

func ordered_pair(a, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

//Q8. TODO: 现在的知识无法解答
//func stack_push(a []int, int) a []int{}

//Q11
func map_it(f func(int) int, a []int) []int {
	var b = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = f(a[i])
	}
	return b
}

//Q13
func bubble_sort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a[:] // ? a是引用类型？？
}

//Q14
func plusTwo() func(int) int {
	return func(in int) int {
		return in + 2
	}
}

func plusX(x int) func(int) int {
	return func(in int) int {
		return in + x
	}
}
