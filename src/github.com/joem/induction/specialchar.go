package main

import "fmt"

type Rect struct {
	width float64
	height float64
}
func main() {
	defer println("system out")
	rect := Rect{10.6, 20.134}
	fmt.Println(rect.size())

}

func (r *Rect) size() float64  {
	fmt.Println(r)
	fmt.Println(&r)
	fmt.Println(*r)
	fmt.Println(&Rect{10, 20})


	fmt.Println("===========================================")
	r = &Rect{123, 234}

	fmt.Println(r)
	fmt.Println(&r)
	fmt.Println(*r)
	fmt.Println(&Rect{10, 20})
	return r.height * r.width
}


