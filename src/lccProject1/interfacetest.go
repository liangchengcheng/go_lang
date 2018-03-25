package main

import (
	"fmt"
)

/* 定义接口 */
type Phone interface {
	call()
}

/* 定义结构体 */
type NokiaPhone struct {
}

/* 实现接口方法 */
func (nokiaPhone NokiaPhone) call() {
	/* 方法实现 */
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	/* 方法实现*/
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
