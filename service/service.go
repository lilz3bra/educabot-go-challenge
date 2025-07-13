package service

import (
	"slices"

	"educabot.com/bookshop/models"
)

func CheapestBook(books []models.Book) (models.Book, bool) {
	if len(books) == 0 {
		return models.Book{}, false // Retorna un libro vacío si no hay libros
	}
	cheapest := slices.MinFunc(books, func(a, b models.Book) int {
		return int(a.Price - b.Price)
	})
	return cheapest, true
}
func BooksWrittenByAuthor(books []models.Book, author string) uint {
	var count uint
	for _, book := range books {
		if book.Author == author {
			count++
		}
	}
	return count
}
func MeanUnitsSold(books []models.Book) uint { // no necesitamos contexto aquí porque no estamos haciendo operaciones asincrónicas
	if len(books) == 0 { // Evita división por cero
		return 0
	}
	var sum uint
	for _, book := range books {
		sum += book.UnitsSold
	}
	return sum / uint(len(books))
}
