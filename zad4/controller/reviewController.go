package controller

import (
	"net/http"

	"zad4/lib"
	"zad4/model"

	"github.com/labstack/echo/v4"
)

type ReviewRequest struct {
	BookID     int    `json:"book_id"`
	UserID     int    `json:"user_id"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}

type reviewController struct{}

func NewReviewController() Controller {
	return &reviewController{}
}

func (a *reviewController) List(ctx echo.Context) error {
	db := lib.GetDB()

	var reviews []model.Review
	if err := db.Preload("User").Preload("Book").Find(&reviews).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string][]model.Review{
		"reviews": reviews,
	})
}

func (a *reviewController) Create(ctx echo.Context) error {
	db := lib.GetDB()

	var reviewRequest ReviewRequest
	if err := ctx.Bind(&reviewRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var book model.Book
	if err := db.First(&book, reviewRequest.BookID).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Book not found"})
	}

	var user model.User
	if err := db.First(&user, reviewRequest.UserID).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "User not found"})
	}

	review := model.Review{
		BookID:     reviewRequest.BookID,
		UserID:     reviewRequest.UserID,
		Rating:     reviewRequest.Rating,
		ReviewText: reviewRequest.ReviewText,
	}

	if err := db.Create(&review).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create review"})
	}

	return ctx.JSON(http.StatusCreated, review)
}

func (a *reviewController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var review model.Review
	if err := db.Preload("User").Preload("Book").First(&review, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Review not found"})
	}

	return ctx.JSON(http.StatusOK, review)
}

func (a *reviewController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var reviewRequest ReviewRequest
	if err := ctx.Bind(&reviewRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var review model.Review
	if err := db.First(&review, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Review not found"})
	}

	var book model.Book
	if err := db.First(&book, reviewRequest.BookID).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Book not found"})
	}

	var user model.User
	if err := db.First(&user, reviewRequest.UserID).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "User not found"})
	}

	review.BookID = reviewRequest.BookID
	review.UserID = reviewRequest.UserID
	review.Rating = reviewRequest.Rating
	review.ReviewText = reviewRequest.ReviewText

	if err := db.Save(&review).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update review"})
	}

	return ctx.JSON(http.StatusOK, review)
}

func (a *reviewController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var review model.Review
	if err := db.First(&review, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Review not found"})
	}

	if err := db.Delete(&review).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete review"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Review deleted successfully"})
}
