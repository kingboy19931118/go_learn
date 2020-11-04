package main

import "fmt"

func main() {
	s1 := []string{"中国", "上海", "浦东新区"}
	var s2 = append(s1, "川沙新镇", "城南路", "园西小区", "34号楼")
	fmt.Println(s1, cap(s1), len(s1))
	fmt.Println(s2, cap(s2), len(s2))

}
