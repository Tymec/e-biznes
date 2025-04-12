package controller

import (
	"net/http"

	"zad4/lib"
	"zad4/model"

	"github.com/labstack/echo/v4"
)

type GenreRequest struct {
	Name string `json:"name"`
}

type genreController struct{}

func NewGenreController() Controller {
	return &genreController{}
}

func (a *genreController) List(ctx echo.Context) error {
	db := lib.GetDB()

	var genres []model.Genre
	if err := db.Find(&genres).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string][]model.Genre{
		"genres": genres,
	})
}

func (a *genreController) Create(ctx echo.Context) error {
	db := lib.GetDB()

	var genreRequest GenreRequest
	if err := ctx.Bind(&genreRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	genre := model.Genre{
		Name: genreRequest.Name,
	}

	if err := db.Create(&genre).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create genre"})
	}

	return ctx.JSON(http.StatusCreated, genre)
}

func (a *genreController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var genre model.Genre
	if err := db.First(&genre, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Genre not found"})
	}

	return ctx.JSON(http.StatusOK, genre)
}

func (a *genreController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var genreRequest GenreRequest
	if err := ctx.Bind(&genreRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var genre model.Genre
	if err := db.First(&genre, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Genre not found"})
	}

	genre.Name = genreRequest.Name

	if err := db.Save(&genre).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update genre"})
	}

	return ctx.JSON(http.StatusOK, genre)
}

func (a *genreController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var genre model.Genre
	if err := db.First(&genre, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Genre not found"})
	}

	var books []model.Book
	if err := db.Where("genre_id = ?", genre.ID).Find(&books).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check genre usage"})
	}
	if len(books) > 0 {
		return ctx.JSON(http.StatusConflict, map[string]string{"error": "Genre is used in books and cannot be deleted"})
	}

	if err := db.Delete(&genre).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete genre"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Genre deleted successfully"})
}
