package route

import (
	"zad4/controller"

	"github.com/labstack/echo/v4"
)

func initControllerRoute(e *echo.Group, ctrl controller.Controller, path string) {
	route := e.Group(path)
	route.GET("", ctrl.List)          // Get all items
	route.POST("", ctrl.Create)       // Create a new item
	route.GET("/:id", ctrl.Get)       // Get an item by ID
	route.PUT("/:id", ctrl.Update)    // Update an item by ID
	route.DELETE("/:id", ctrl.Delete) // Delete an item by ID
}

func InitApiRoutes(e *echo.Echo) {
	apiRoute := e.Group("/api")
	initControllerRoute(apiRoute, controller.NewBookController(), "/books")
	initControllerRoute(apiRoute, controller.NewAuthorController(), "/authors")
	initControllerRoute(apiRoute, controller.NewGenreController(), "/genres")
	initControllerRoute(apiRoute, controller.NewReviewController(), "/reviews")
	initControllerRoute(apiRoute, controller.NewUserController(), "/users")
}
