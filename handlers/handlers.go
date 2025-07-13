package handlers

import (
	"net/http"

	"educabot.com/bookshop/providers"
	"educabot.com/bookshop/service"
	"github.com/gin-gonic/gin"
)

type GetMetricsRequest struct {
	Author string `form:"author"`
}

func NewGetMetrics(booksProvider providers.BooksProvider) GetMetrics {
	return GetMetrics{booksProvider}
}

type GetMetrics struct {
	booksProvider providers.BooksProvider
}

func (h GetMetrics) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var query GetMetricsRequest
		if err := ctx.ShouldBindQuery(&query); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "query invalida"})
			return
		}
		books, err := h.booksProvider.GetBooks(ctx) // usar el contexto de gin en vez de uno nuevo
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get books"})
			return
		}

		meanUnitsSold := service.MeanUnitsSold(books)
		cheapest, found := service.CheapestBook(books)
		var cheapestBook string
		if found {
			cheapestBook = cheapest.Name
		} else {
			cheapestBook = ""
		}
		booksWrittenByAuthor := service.BooksWrittenByAuthor(books, query.Author)

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"mean_units_sold":         meanUnitsSold,
			"cheapest_book":           cheapestBook,
			"books_written_by_author": booksWrittenByAuthor,
		})
	}
}
