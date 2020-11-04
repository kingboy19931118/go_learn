package main

import "fmt"

func main() {
	defer func() {
		if err:=recover();err!= nil {
			fmt.Println("father 捕获panic代码异常")
			fmt.Println("father 通过 recover 处理进行恢复")
		}
	}()
	GO()
	PHP()
	JAVA()
}

func GO() {
	fmt.Println("I`m golang")
}

func PHP() {
	defer func() {
		if err:=recover();err!= nil {
			fmt.Println("php 捕获panic代码异常")
			fmt.Println("php 通过 recover 处理进行恢复")
		}
	}()
	panic("PHP 抛出Panic")
	fmt.Println("I`m PHP")
}

func JAVA() {
	fmt.Println("I`m JAVA")
}