package main

import (
	"fmt"
	"net/http"
)

func main() {
	count := 0
	// 条件表达式没有括号
	// 支持一个初始化表达式 （可以是多变量初始化语句）
	// 左大括号必须和条件语句同一行
	if number:=7; number <1 {
		fmt.Println(1)
	} else if number>=1 && number<= 10 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}

	http.HandleFunc("/count", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte ("总访问次数：" + string(count)))
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if err1 := request.ParseForm() ; err1 != nil {
			fmt.Println(err1)
			writer.Write([]byte ("系统异常"))
		}
		writer.Write([]byte ("hello world  : " + request.Form.Get("name")))
	})
	// 优雅处理err
	if err:= http.ListenAndServe(":9090", nil) ; err != nil {
		fmt.Println(err)
	}
}
