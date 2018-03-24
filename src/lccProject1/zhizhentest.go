package main

import "fmt"

func main() {
	// 声明实际的变量
	var a int = 20
	// 声明指针变量
	var ip *int
	// 指针变量的存储地址
	ip = &a;
	fmt.Println("a 变量的地址是", &a)
	fmt.Println("ip变量储存指针地址", ip)
	fmt.Println("* ip变量的值", *ip)
}
