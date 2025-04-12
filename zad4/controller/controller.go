package controller

import (
	"github.com/labstack/echo/v4"
)

type Controller interface {
	List(ctx echo.Context) error
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
