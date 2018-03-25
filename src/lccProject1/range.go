package main

import "fmt"

func main() {
	// 这里我们使用range去求一个slice 和 使用数组跟这个很类似
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// 在数组上使用range 将传入的index和值2个变量
	// 上面的例子我们不需要使用元素的序号
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index :", i)
		}
	}

	// range 也可以使用在map上面
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s - > %s\n", k, v)
	}

	// range 也可以用来unicode字符串
	// 第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i ,c := range "go" {
		fmt.Println(i, c)
	}
}
