package main

var x, y int

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

// 不带声明格式的只能在函数体内出现
// g, h := 123, "hello world"

func main() {
	g, h := 123, "hello"
	println(g, h, a, b, c, d, e, f, x, y)

}

// 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：

// 当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
