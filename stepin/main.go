// stepin project main.go
package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func CreateFile(fd int, name string) *os.File {
	if fd < 0 {
		return nil
	}
	f, _ := os.Create(name)

	return f
}

type NameAge struct {
	name string
	age  int
}

func SetItsName(n1 *NameAge, n string) {
	if n1 != nil {
		n1.name = n
	}
}

func (n1 *NameAge) SetName(n string) {
	if n1 != nil {
		n1.name = n
	}
}

func (n1 *NameAge) SetAge(a int) {
	if n1 != nil {
		n1.age = a
	}
}

func (n1 *NameAge) GetName() string {
	return n1.name
}

func (n1 *NameAge) GetAge() int {
	return n1.age
}

type I interface {
	GetName() string
	SetName(string)
}

func fi(p I) {
	p.SetName("function op on Interface")
	fmt.Println(p.GetName())
}

func (v1 NameAge) SetEmpty() {
	v1.age = 18
	fmt.Println(v1)
}

type NameAge2 struct {
	name string
}

func (n1 *NameAge2) SetName(n string) {
	if n1 != nil {
		n1.name = n
	}
}

func main() {
	fmt.Println("Hello World!")

	var p *int
	if p == nil {
		p = new(int)
	}
	*p = 8
	fmt.Printf("*%v = %v\n", p, *p)

	var q *SyncedBuffer = new(SyncedBuffer)
	fmt.Printf("*%v = %v\n", q, *q)

	var r SyncedBuffer
	fmt.Printf("*%v = %v\n", &r, r)

	a := new(NameAge)
	a.name = "oney"
	a.age = 40
	fmt.Println(a)

	var b NameAge
	b.SetName("billy")
	(&b).SetAge(31)
	b.SetEmpty()
	(&b).SetEmpty()
	fmt.Println(b)

	var pb *NameAge = &b
	pb.SetName("pbNo")
	(*pb).SetAge(32)
	pb.SetEmpty()
	fmt.Println(pb)

	fi(&b)

	type NewNameAge NameAge
	type pNewNameAge *NameAge //*NewNameAge
	type PrintableNameAge struct{ NameAge }

	// 通过type定义出来的东西很怪：
	// A type declaration binds an identifier, the type name, to a new type that has the same underlying type as an existing type, and operations defined for the existing type are also defined for the new type. The new type is different from the existing type.
	// The declared type does not inherit any methods bound to the existing type, but the method set of an interface type or of elements of a composite type remains unchanged.
	// 1.无论申明在类型还是类型指针（这两个在接口层面）上的方法，通过type定义后都丢失了！
	// 2.如果非要保留，则只能通过组合的方式定义新类型。

	//示例一：
	//var c NewNameAge
	//c.SetName("NewType")  // c.SetName undefined
	//c.SetEmpty()	//

	//示例二：
	//var pc pNewNameAge = &b
	//pc.SetName("pNewType")

	//示例三：
	var d PrintableNameAge
	d.SetName("Printable")

	//示例四：结构中的两个子结构拥有同样的方法名！怎么调用？
	type PrintableNameAge2 struct {
		NameAge
		NameAge2
	}
	var e PrintableNameAge2
	//e.SetName("which one") //ambiguous，有两个同名方法，需要
	e.SetAge(100)
	e.SetEmpty()
	e.NameAge.SetName("我是Set1，你猜对了吗？")
	e.NameAge2.SetName("我TMD是Set2！")
	fmt.Println(e)

	// 基本类型转换：除了float可以被转为int/uint，其他都符合直觉
	byteslice := []byte(e.NameAge.name)
	runeslice := []rune(e.NameAge.name)
	fmt.Println(byteslice)
	fmt.Println(string(byteslice))
	fmt.Println(runeslice)
	fmt.Println(string(runeslice))
	var ifl uint32 = 189
	fmt.Printf("uint32: %v, to float32: %v, to float64: %v\n",
		ifl, float32(ifl), float64(ifl))
	var fil float64 = -199.8
	fmt.Printf("float32: %v, to int: %v, to uint32: %v, to uint64: %v\n",
		fil, int(fil), uint32(fil), uint64(fil))

	// 自定义类型类型转换
	var yoo = NewNameAge{"haha", 12}
	var foo = NameAge(yoo)
	fmt.Println(foo)
	//yoo.SetName("yoo") 是不行的！
	//--有点意思！就是type定义出来的没有接口！
	foo.SetName("foo") //类型转换后，原先没有的也有了接口！
	fmt.Println(foo)

}
