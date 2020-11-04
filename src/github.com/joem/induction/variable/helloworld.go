package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	println("hello world 222")

	const number1 = 10

	var str string
	var number2 int
	var number3 [5] int
	var number4 [] int
	var number5 map [string] int
	type number6 int
	type number7 struct {
		fmt.Formatter
	}

	type number8 interface {

	}

	var name = "张三"
	name = "里斯"
	const sex = 'Y'
	println(name, sex)
	println(number1 , number2)
	println(number5)
	println(str)
	println("number3: ",len(number3), " number4: ", len(number4), "number5", len(number5))

	var b, c string = "b", "c"
	fmt.Println(b, c)

	d := true
	println(d)

	var d1 = true
	println(d1)

	var i int8 = -128
	var fl float64
	fmt.Println("int8 type i : " , i)
	fmt.Println(fl)

}

func demo1 () {

}