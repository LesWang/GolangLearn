package main

import (
    "fmt"
    "github.com/shawnfeng/sutil/slog"
    )

type Books struct {
    title string
    author string
    subject string
    bookId int
}

func main()  {
    slog.Warnln("Hello")
    var book1 Books = Books{"Go语言", "www.runoob.com", "Go语言教程", 123456}
    fmt.Println(book1)
    var book2 = Books{title: "Go语言", author: "www.runoob.com", subject: "Go语言教程", bookId: 123456}
    fmt.Println(book2)
    var book3 = Books{title: "Go语言", author: "www.runoob.com"}
    fmt.Println(book3)

    var book4 Books
    book4.title = "Python语言"
    book4.author = "www.runoob.com"
    book4.subject = "Python语言教程"
    book4.bookId = 123457
    fmt.Println(book4.title, book4.author)

    var bookPtr = &book4
    fmt.Println(bookPtr.title, bookPtr.author)
}
