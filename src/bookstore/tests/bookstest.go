package tests

import (
	"strings"
	"github.com/robfig/revel"
)

type BooksTest struct {
	revel.TestSuite
}

func (t BooksTest) TestThatIndexPageWorks() {
	t.Get("/books")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t BooksTest) TestThatPostingToBooksControllerWorks() {
	data := strings.NewReader(`{"title":"fubar"}`)
	t.Post("/books", "application/json", data)

	t.AssertOk()
	t.AssertContentType("application/json")
}
