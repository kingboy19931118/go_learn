package main

import "fmt"

func main() {

	x := make([]int, 0, 5)
	fmt.Printf("value:%v, len:%d, cap:%d", x, len(x), cap(x))
	fmt.Println()

	//x = append(x, 0, 1,2,3,4,5,6,7,8,9,10,11)
	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Printf("value:%v, len:%d, cap:%d", x, len(x), cap(x))
	fmt.Println()
	fmt.Println(x)

	m := make(map[string]int)
	m["a"] = 1
	x2, ok := m["b"]
	fmt.Println(x2, ok)

	delete(m, "a")

	x21, ok := m["a"]
	fmt.Println(x21, ok)

	var m3 manager
	m3.name = "Tom"
	m3.age = 29
	m3.title = "CTO"
	fmt.Println(m3)

	var x4 X
	x4.inc()
	println(x4)

	fmt.Println(m3.ToString())

	var p Printer = m3
	p.Print()

	var o Object = m3
	fmt.Println(o)
}


type X int

func (x *X) inc() {
	*x++
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}


type user struct {
	name string
	age byte
}
type manager struct {
	user
	title string
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface {
	ToString() string
	Print()
}

type Object interface {

}



