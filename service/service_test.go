package service

import (
	"testing"

	"educabot.com/bookshop/models"
	"github.com/stretchr/testify/assert"
)

func TestMeanUnitsSold(t *testing.T) {
	// ARRANGE: Define test cases
	testCases := []struct {
		name     string
		books    []models.Book
		expected uint
	}{
		{
			name: "Test with multiple books",
			books: []models.Book{
				{Name: "Book A", UnitsSold: 100},
				{Name: "Book B", UnitsSold: 200},
				{Name: "Book C", UnitsSold: 300},
			},
			expected: 200, // (100 + 200 + 300) / 3
		},
		{
			name: "Test with a single book",
			books: []models.Book{
				{Name: "Book A", UnitsSold: 50},
			},
			expected: 50,
		},
		{
			name:     "Test with an empty slice",
			books:    []models.Book{},
			expected: 0, // Should not panic
		},
		{
			name: "Test with zero units sold",
			books: []models.Book{
				{Name: "Book A", UnitsSold: 0},
				{Name: "Book B", UnitsSold: 0},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ACT: Call the function being tested
			result := MeanUnitsSold(tc.books)

			// ASSERT: Check if the result is what we expect
			assert.Equal(t, tc.expected, result, "The calculated mean should match the expected value.")
		})
	}
}

func TestCheapestBook(t *testing.T) {
	// ARRANGE: Define our books and test cases
	book1 := models.Book{Name: "Book A", Price: 20}
	book2 := models.Book{Name: "Book B", Price: 10} // Cheapest
	book3 := models.Book{Name: "Book C", Price: 30}

	testCases := []struct {
		name          string
		books         []models.Book
		expectedBook  models.Book
		expectedFound bool
	}{
		{
			name:          "Test with multiple books",
			books:         []models.Book{book1, book2, book3},
			expectedBook:  book2,
			expectedFound: true,
		},
		{
			name:          "Test with a single book",
			books:         []models.Book{book3},
			expectedBook:  book3,
			expectedFound: true,
		},
		{
			name:          "Test with an empty slice",
			books:         []models.Book{},
			expectedBook:  models.Book{}, // Expect the zero-value for Book
			expectedFound: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ACT: Call the function
			result, found := CheapestBook(tc.books)

			// ASSERT: Check the results
			assert.Equal(t, tc.expectedFound, found)
			assert.Equal(t, tc.expectedBook, result, "The returned book should be the cheapest one.")
		})
	}
}

func TestBooksWrittenByAuthor(t *testing.T) {
	// ARRANGE: Define our books and test cases
	books := []models.Book{
		{Author: "J.K. Rowling"},
		{Author: "J.R.R. Tolkien"},
		{Author: "J.K. Rowling"},
		{Author: "George R.R. Martin"},
	}

	testCases := []struct {
		name          string
		authorToFind  string
		books         []models.Book
		expectedCount uint
	}{
		{
			name:          "Test for author with multiple books",
			authorToFind:  "J.K. Rowling",
			books:         books,
			expectedCount: 2,
		},
		{
			name:          "Test for author with one book",
			authorToFind:  "J.R.R. Tolkien",
			books:         books,
			expectedCount: 1,
		},
		{
			name:          "Test for author with no books",
			authorToFind:  "Stephen King",
			books:         books,
			expectedCount: 0,
		},
		{
			name:          "Test with an empty book slice",
			authorToFind:  "J.K. Rowling",
			books:         []models.Book{},
			expectedCount: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ACT: Call the function
			result := BooksWrittenByAuthor(tc.books, tc.authorToFind)

			// ASSERT: Check the result
			assert.Equal(t, tc.expectedCount, result, "The count of books by the author should match the expected value.")
		})
	}
}
