package api

import (
	"encoding/json"
	"net/http"
)

// Book type with Name, Author and ISBN
type Book struct {
	// define the book
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var books = map[string]Book{
	"1569319200": Book{
		Title:       "Dragon Ball",
		Author:      "Akira Toriyama",
		ISBN:        "1569319200",
		Description: "a Japanese media franchise created in 1984",
	},
	"1569319049": Book{
		Title:  "Yu Yu Hakusho",
		Author: "Yoshihiro Togashi",
		ISBN:   "1569319049",
	},
}

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}

	return book
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
	v := make([]Book, 0, len(books))
	for _, value := range books {
		v = append(v, value)
	}

	return v
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func writeJSON(w http.ResponseWriter, books []Book) {
	b, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
