package main

import "fmt"

func main()  {
	var a int  =10
	fmt.Println("变量的地址是",&a)
}

// 什么是指针呢？
// 指针是可以指向任何一个值的内存地址它指向那个值的内存地址
// 类似变量和常量，在使用指针之前你需要声明指针，

// 指向整形
var ip  *int
// 指向浮点型
var fp *float32