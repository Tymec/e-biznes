package controller

import (
	"fmt"
	"net/http"

	"zad4/lib"
	"zad4/model"

	"github.com/labstack/echo/v4"
)

type BookRequest struct {
	Title     string `json:"title"`
	Year      int    `json:"year"`
	ISBN      string `json:"isbn"`
	Pages     int    `json:"pages"`
	AuthorIds []int  `json:"author_ids"`
	GenreIds  []int  `json:"genre_ids"`
}

type bookController struct{}

func NewBookController() Controller {
	return &bookController{}
}

func (a *bookController) List(ctx echo.Context) error {
	db := lib.GetDB()

	var books []model.Book
	if err := db.Model(&model.Book{}).Preload("Authors").Preload("Genres").Find(&books).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string][]model.Book{
		"books": books,
	})
}

func (a *bookController) Create(ctx echo.Context) error {
	db := lib.GetDB()

	var bookRequest BookRequest
	if err := ctx.Bind(&bookRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var authors []model.Author
	for _, authorID := range bookRequest.AuthorIds {
		var author model.Author
		if err := db.First(&author, authorID).Error; err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Author with ID %d not found", authorID)})
		}
		authors = append(authors, author)
	}

	var genres []model.Genre
	for _, genreID := range bookRequest.GenreIds {
		var genre model.Genre
		if err := db.First(&genre, genreID).Error; err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Genre with ID %d not found", genreID)})
		}
		genres = append(genres, genre)
	}

	book := model.Book{
		Title:   bookRequest.Title,
		Year:    bookRequest.Year,
		ISBN:    bookRequest.ISBN,
		Pages:   bookRequest.Pages,
		Authors: authors,
		Genres:  genres,
	}

	if err := db.Create(&book).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create book"})
	}

	return ctx.JSON(http.StatusCreated, book)
}

func (a *bookController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var book model.Book
	if err := db.Preload("Authors").Preload("Genres").First(&book, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
	}

	return ctx.JSON(http.StatusOK, book)
}

func (a *bookController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var bookRequest BookRequest
	if err := ctx.Bind(&bookRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
	}

	var authors []model.Author
	for _, authorID := range bookRequest.AuthorIds {
		var author model.Author
		if err := db.First(&author, authorID).Error; err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Author with ID %d not found", authorID)})
		}
		authors = append(authors, author)
	}

	var genres []model.Genre
	for _, genreID := range bookRequest.GenreIds {
		var genre model.Genre
		if err := db.First(&genre, genreID).Error; err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("Genre with ID %d not found", genreID)})
		}
		genres = append(genres, genre)
	}

	book.Title = bookRequest.Title
	book.Year = bookRequest.Year
	book.ISBN = bookRequest.ISBN
	book.Pages = bookRequest.Pages
	book.Authors = authors
	book.Genres = genres

	if err := db.Save(&book).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update book"})
	}

	return ctx.JSON(http.StatusOK, book)
}

func (a *bookController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
	}

	var reviews []model.Review
	if err := db.Where("book_id = ?", book.ID).Find(&reviews).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to find reviews"})
	}
	for _, review := range reviews {
		if err := db.Delete(&review).Error; err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete review"})
		}
	}

	if err := db.Delete(&book).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete book"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}
