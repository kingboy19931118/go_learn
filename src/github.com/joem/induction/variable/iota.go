package main

const(
GB int = 1 << (iota * 10)
MB int = 1 << (iota * 10)
KB int = 1 << (iota * 10)
c = iota
)

const (
	a int = iota
	b int = iota
	b1
)

func main() {
	println(GB, MB, KB, c)
	println (a, b, b1)
	var k , _ int  = demo()

	println(k)

}

func demo() (a, b int) {
	return 1024, 2048
}
