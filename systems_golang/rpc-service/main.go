package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var (
	ErrISBN      = errors.New("missing ISBN")
	ErrDuplicate = errors.New("duplicate book")
	ErrMissing   = errors.New("missing book")
)

type Book struct {
	ISBN          string
	Title, Author string
	Year, Pages   int
}

type ReadingList struct {
	Books    []Book
	Progress []int
}

// get the index of the book using its isbn
func (r *ReadingList) bookIndex(isbn string) int {
	for i := range r.Books {
		if isbn == r.Books[i].ISBN {
			return i
		}
	}
	return -1
}

// AddBook adds a unique book to list
func (r *ReadingList) AddBook(b Book) error {
	if b.ISBN == "" {
		return ErrISBN
	}
	if r.bookIndex(b.ISBN) != -1 {
		return ErrDuplicate
	}
	r.Books = append(r.Books, b)
	r.Progress = append(r.Progress, 0)
	return nil
}

// RemoveBook from reading list
func (r *ReadingList) RemoveBook(isbn string) error {
	if isbn == "" {
		return ErrISBN
	}
	i := r.bookIndex(isbn)
	if i == -1 {
		return ErrMissing
	}

	// replace deleted book with last element of list
	r.Books[i] = r.Books[len(r.Books)-1]
	r.Progress[i] = r.Progress[len(r.Progress)-1]

	// shrink the list to remove last element
	r.Books = r.Books[:len(r.Books)-1]
	r.Progress = r.Progress[:len(r.Progress)-1]
	return nil
}

// GetProgress
func (r *ReadingList) GetProgress(isbn string) (int, error) {
	if isbn == "" {
		return -1, ErrISBN
	}
	i := r.bookIndex(isbn)
	if i == -1 {
		return -1, ErrMissing
	}
	return r.Progress[i], nil
}

// SetProgress
func (r *ReadingList) SetProgress(isbn string, pages int) error {
	if isbn == "" {
		return ErrISBN
	}
	i := r.bookIndex(isbn)
	if p := r.Books[i].Pages; pages > p {
		pages = p
	}
	r.Progress[i] = pages
	return nil
}

// AdvanceProgress
func (r *ReadingList) AdvanceProgress(isbn string, pages int) error {
	if isbn == "" {
		return ErrMissing
	}
	i := r.bookIndex(isbn)
	if i == -1 {
		return ErrMissing
	}
	if p := r.Books[i].Pages - r.Progress[i]; p < pages {
		pages = p
	}
	r.Progress[i] += pages
	return nil
}

// ReadingService is an rpc service for reading list
type ReadingService struct{ ReadingList }

// sets the success pointer value from error
func setSucess(err error, b *bool) error {
	*b = err == nil
	return err
}

func (r *ReadingService) AddBook(b Book, success *bool) error {
	return setSucess(r.ReadingList.AddBook(b), success)
}

func (r *ReadingService) RemoveBook(isbn string, success *bool) error {
	return setSucess(r.ReadingList.RemoveBook(isbn), success)
}

func (r *ReadingService) GetProgress(isbn string, pages *int) (err error) {
	*pages, err = r.ReadingList.GetProgress(isbn)
	return
}

// Progress is a struct because we can only pass two arguments to the rpc
// service methods of which the second must be a pointer
type Progress struct {
	ISBN  string
	Pages int
}

func (r *ReadingService) SetProgress(p Progress, success *bool) error {
	return setSucess(r.ReadingList.SetProgress(p.ISBN, p.Pages), success)
}

func (r *ReadingService) AdvanceProgress(p Progress, success *bool) error {
	return setSucess(r.ReadingList.AdvanceProgress(p.ISBN, p.Pages), success)
}

func main() {
	bookService := new(ReadingService)
	if err := rpc.Register(bookService); err != nil {
		log.Fatalln(err)
	}
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server Started")
	if err := http.Serve(l, nil); err != nil {
		log.Fatal(err)
	}
}
