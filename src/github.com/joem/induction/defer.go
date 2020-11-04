package main

func main() {
	println("demo1 execute: " , demo1())
	println("demo2 execute: " , demo2())
	println("demo3 execute: " , *demo3())
}

func demo3() *int {
	var i int
	defer func() {
		i++
		println("defer 2 : ", i)
	}()

	defer func() {
		i++
		println("defer 1 : ", i)
	}()
	return &i
}

func demo1() (i int) {
	defer func() {
		i++
		println("defer 2 : ", i)
	}()

	defer func() {
		i++
		println("defer 1 : ", i)
	}()
	return
}

func demo2() int {
	i := 0
	defer func() {
		i++
		println("defer 2 : ", i)
	}()
	defer func() {
		i++
		println("defer 2 : ", i)
	}()
	return i
}

