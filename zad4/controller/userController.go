package controller

import (
	"net/http"

	"zad4/lib"
	"zad4/model"

	"github.com/labstack/echo/v4"
)

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userController struct{}

func NewUserController() Controller {
	return &userController{}
}

func (a *userController) List(ctx echo.Context) error {
	db := lib.GetDB()

	var users []model.User
	if err := db.Omit("Reviews").Find(&users).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string][]model.User{
		"users": users,
	})
}

func (a *userController) Create(ctx echo.Context) error {
	db := lib.GetDB()

	var userRequest UserRequest
	if err := ctx.Bind(&userRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user := model.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	if err := db.Create(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (a *userController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var user model.User
	if err := db.Preload("Reviews").First(&user, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return ctx.JSON(http.StatusOK, user)
}

func (a *userController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var userRequest UserRequest
	if err := ctx.Bind(&userRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	user.Name = userRequest.Name
	user.Email = userRequest.Email

	if err := db.Save(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	return ctx.JSON(http.StatusOK, user)
}

func (a *userController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	db := lib.GetDB()

	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	var reviews []model.Review
	if err := db.Where("user_id = ?", user.ID).Find(&reviews).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check user reviews"})
	}
	if len(reviews) > 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "User has reviews and cannot be deleted"})
	}

	if err := db.Delete(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}

	return ctx.NoContent(http.StatusNoContent)
}
