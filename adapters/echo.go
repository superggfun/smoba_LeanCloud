package adapters

import (
	"github.com/labstack/echo/v4"
	"github.com/leancloud/go-sdk/leancloud"
)

func Echo(e *echo.Echo) {
	e.Any("/1/*", echo.WrapHandler(leancloud.Engine.Handler()), setResponseContentType)
	e.Any("/1.1/*", echo.WrapHandler(leancloud.Engine.Handler()), setResponseContentType)
	e.Any("/__engine/*", echo.WrapHandler(leancloud.Engine.Handler()), setResponseContentType)
}

func setResponseContentType(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json; charset=UTF-8")
		return next(c)
	}
}
