package main

import (
	"eatobin.com/totalbeginnergo/book"
	"eatobin.com/totalbeginnergo/borrower"
	"eatobin.com/totalbeginnergo/library"
	"fmt"
	"io/ioutil"
)

var borrowers []borrower.Borrower
var books []book.Book

func main() {
	borrowers = library.AddBorrower(borrowers, borrower.NewBorrower("Jim", 3))
	borrowers = library.AddBorrower(borrowers, borrower.NewBorrower("Sue", 3))
	books = library.AddBook(books, book.NewBook("War And Peace", "Tolstoy"))
	books = library.AddBook(books, book.NewBook("Great Expectations", "Dickens"))
	fmt.Println("\nJust created new library")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Check out War And Peace to Sue")
	books = library.CheckOut("Sue", "War And Peace", borrowers, books)
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Now check in War And Peace from Sue...")
	books = library.CheckIn("War And Peace", books)
	fmt.Println("...and check out Great Expectations to Jim")
	books = library.CheckOut("Jim", "Great Expectations", borrowers, books)
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Add Eric and The Cat In The Hat")
	borrowers = library.AddBorrower(borrowers, borrower.Borrower{Name: "Eric", MaxBooks: 1})
	books = library.AddBook(books, book.Book{Title: "The Cat In The Hat", Author: "Dr. Seuss"})
	fmt.Println("Check Out Dr. Seuss to Eric")
	books = library.CheckOut("Eric", "The Cat In The Hat", borrowers, books)
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Now let's do some BAD stuff...")

	fmt.Println("Add a borrower that already exists (total.Borrower('Jim', 3))")
	borrowers = library.AddBorrower(borrowers, borrower.Borrower{Name: "Jim", MaxBooks: 3})
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Add a book that already exists (total.Book('War And Peace', 'Tolstoy', None))")
	books = library.AddBook(books, book.Book{Title: "War And Peace", Author: "Tolstoy"})
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Check out a valid book to an invalid person (checkOut('JoJo', 'War And Peace', borrowers))")
	books = library.CheckOut("JoJo", "War And Peace", borrowers, books)
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Check out an invalid book to an valid person (checkOut('Sue', 'Not A total.Book', borrowers))")
	books = library.CheckOut("Sue", "Not A Book", borrowers, books)
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Last - check in a book not checked out (checkIn('War And Peace'))")
	books = library.CheckIn("War And Peace", books)
	fmt.Println("No change to Test Library:")
	fmt.Println(library.StatusToString(books, borrowers))

	fmt.Println("Okay... let's finish with some persistence. First clear the whole library:")
	newEmptyV()

	fmt.Println("Lets read in a new library from \"borrowers-before.json\" and \"books-before.json\":")
	//newV(tvBooks, tvBorrowers, jsonBorrowersFileBefore, jsonBooksFile)
	//fmt.Println("Add... a new borrower:")
	//tvBorrowers.transform(addItem(Borrower("BorrowerNew", 300), _))
	//fmt.Println(statusToString(tvBooks.get, tvBorrowers.get))
}

func newEmptyV() {
	borrowers = library.ZeroBorrowers
	books = library.ZeroBooks
	fmt.Println(library.StatusToString(books, borrowers))
}

func ReadFileIntoJsonString(fp string) (string, error) {
	jsonStr, err := ioutil.ReadFile(fp)
	return string(jsonStr), err
}

func NewV(brsfp string, bksfp string) ([]borrower.Borrower, []book.Book) {
	jsonBrsStr, brFileErr := ReadFileIntoJsonString(brsfp)
	if brFileErr != nil {
		fmt.Println(brFileErr.Error())
		return library.ZeroBorrowers, library.ZeroBooks
	}
	jsonBksStr, bkFileErr := ReadFileIntoJsonString(bksfp)
	if bkFileErr != nil {
		fmt.Println(bkFileErr.Error())
		return library.ZeroBorrowers, library.ZeroBooks
	}
	brs, brParseErr := library.JsonStringToBorrowers(jsonBrsStr)
	if brParseErr != nil {
		fmt.Println(brParseErr.Error())
		return library.ZeroBorrowers, library.ZeroBooks
	}
	bks, bkParseErr := library.JsonStringToBooks(jsonBksStr)
	if bkParseErr != nil {
		fmt.Println(bkParseErr.Error())
		return library.ZeroBorrowers, library.ZeroBooks
	}
	return brs, bks
}
