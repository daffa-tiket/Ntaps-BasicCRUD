package infrastructure

import (
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func RegisterMiddleware(holder *Holder) {
	holder.SharedHolder.Echo.Use(middleware.Recover())

	// - register your custom middleware here
	// - holder.SharedHolder.Echo.Use(HelloWorldMiddleware(holder))
}

func HelloWorldMiddleware(h *Holder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			h.SharedHolder.Logger.Info("hello world")
			return next(ctx)
		}
	}
}
