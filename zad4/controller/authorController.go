package controller

import (
	"net/http"

	"zad4/lib"
	"zad4/model"

	"github.com/labstack/echo/v4"
)

type AuthorRequest struct {
	Name string `json:"name"`
}

type authorController struct{}

func NewAuthorController() Controller {
	return &authorController{}
}

func (a *authorController) List(ctx echo.Context) error {
	db := lib.GetDB()

	var authors []model.Author
	if err := db.Find(&authors).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string][]model.Author{
		"authors": authors,
	})
}

func (a *authorController) Create(ctx echo.Context) error {
	db := lib.GetDB()

	var authorRequest AuthorRequest
	if err := ctx.Bind(&authorRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	author := model.Author{
		Name: authorRequest.Name,
	}

	if err := db.Create(&author).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create author"})
	}

	return ctx.JSON(http.StatusCreated, author)
}

func (a *authorController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var author model.Author
	if err := db.First(&author, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Author not found"})
	}

	return ctx.JSON(http.StatusOK, author)
}

func (a *authorController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var authorRequest AuthorRequest
	if err := ctx.Bind(&authorRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var author model.Author
	if err := db.First(&author, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Author not found"})
	}

	author.Name = authorRequest.Name

	if err := db.Save(&author).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update author"})
	}

	return ctx.JSON(http.StatusOK, author)
}

func (a *authorController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var author model.Author
	if err := db.First(&author, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Author not found"})
	}

	var books []model.Book
	if err := db.Where("author_id = ?", id).Find(&books).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check books"})
	}
	if len(books) > 0 {
		return ctx.JSON(http.StatusConflict, map[string]string{"error": "Author cannot be deleted because they have books associated with them"})
	}

	if err := db.Delete(&author).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete author"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Author deleted successfully"})
}
