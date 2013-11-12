package tests

import (
	"encoding/json"
	"strings"
	"github.com/stretchr/testify/assert"
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
	body := t.ResponseBody
	var jsonBody map[string]interface{}
	_ = json.Unmarshal(body, &jsonBody)
	t.Assert(assert.ObjectsAreEqual(jsonBody, map[string]interface{} { "title" : "fubar" }))
}
