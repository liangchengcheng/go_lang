package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main()  {
	var Book1 Books
	var Book2 Books

	// book1 的描述
	Book1.title = "go 语言学习"
	Book1.author = "www.go.com"
	Book1.subject= "go 语言教程"
	Book1.book_id = 645426

	// book 2 描述
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	// 打印book的信息
	printBook(Book1)
	printBook(Book2)
}

func printBook(book Books)  {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}

