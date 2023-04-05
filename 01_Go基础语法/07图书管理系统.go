package main

import "fmt"

type Book struct {
	id     int
	name   string
	author string
}

type BookManager struct {
	books []Book
}

// AddBook 添加图书
func (bm *BookManager) AddBook(book Book) {
	bm.books = append(bm.books, book)
}

// DeleteBook 删除图书
func (bm *BookManager) DeleteBook(id int) bool {
	for i, book := range bm.books {
		if book.id == id {
			bm.books = append(bm.books[:i], bm.books[i+1:]...)
			return true
		}
	}
	return false
}

// FindBook 查找图书
func (bm *BookManager) FindBook(id int) *Book {
	for i, book := range bm.books {
		if book.id == id {
			return &bm.books[i]
		}
	}
	return nil
}

// ViewBooks 展示图书
func (bm *BookManager) ViewBooks() {
	fmt.Println("Books List:")
	for _, book := range bm.books {
		fmt.Printf("ID:%d, name:%s, author:%s\n", book.id, book.name, book.author)
	}
}

func demo007() {
	// 创建图书管理器
	// bm := BookManager{
	// 	 books: []Book{},
	// }
	// bm := new(BookManager)
	// bm := BookManager{}
	bm := &BookManager{}

	// 添加图书
	bm.AddBook(Book{id: 1, name: "The Great Gatsby", author: "F. Scott Fitzgerald"})
	bm.AddBook(Book{id: 2, name: "To Kill a Mockingbird", author: "Harper Lee"})
	bm.AddBook(Book{id: 3, name: "1984", author: "George Orwell"})

	// 显示图书列表
	bm.ViewBooks()

	// 查找图书
	fmt.Println("Find Book:")
	book := bm.FindBook(2)
	if book != nil {
		fmt.Printf("ID: %d, name: %s, Author: %s\n", book.id, book.name, book.author)
	} else {
		fmt.Println("Book not found.")
	}

	// 删除图书
	fmt.Println("Delete Book:")
	if bm.DeleteBook(3) {
		bm.ViewBooks()
	} else {
		fmt.Println("Book not found.")
	}
}
