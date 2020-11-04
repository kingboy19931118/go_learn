package main

import "fmt"

const (
	am1      int = 1
	dm1      int = 0
	bm1, cm1     = 1, 1
)

func main() {
	println(am1, dm1, bm1, cm1)

	// type byte = unit8 不带符号8位
	// type rune = int32 带符号32位
	str := "Hello!一二三"

	strList1 := []byte(str)
	strList2 := []rune(str)

	// go 中字符串底层默认是 [] byte 形式存储的
	fmt.Println(len(str))
	fmt.Printf("byte:%s---%d---%b\n", str, len(str), []byte(str))
	// UTF-8编码：一个中文字符等于三个字节，显然一个unit8无法存储三个字节
	fmt.Println(strList1)
	fmt.Printf("byte:%s---%d---%b\n", string(strList1), len(strList1), strList1)
	fmt.Println(strList2)
	fmt.Printf("rune:%s---%d---%b\n", string(strList2), len(strList2), strList2)

	// 循环取值
	for i := 0; i < len(str); i++ {
		// 乱码
		fmt.Print(string(str[i]))
	}

	fmt.Println()

	for i := 0; i < len(strList2); i++ {
		fmt.Print(string(strList2[i]))
	}

	fmt.Println()

	for _, v := range str {
		fmt.Print(string(v))
	}

}
