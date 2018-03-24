package main

import "fmt"

// var + name + [size] + 类型
var balance [10] float32

// 数组的初始化
var balances = [5]float32{10000.2,2.0,3.4}
// 初始化数组中 大括号的长度不能大于中括号的长度

func main()  {
	var n[10] int
	var i, j int
	for i = 0;i < 10 ;i++  {
		n[i] = i +100
	}

	for j = 0; j < 10 ; j++ {
		fmt.Println("element " ,j,n[j])

	}
}

