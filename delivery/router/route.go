package router

import (
	"context"
	"sirclo/delivery/middlewares"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, srv *handler.Server) {
	e.Use(middleware.CORSWithConfig((middleware.CORSConfig{})))

	e.POST("/query", func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "EchoContextKey", c.Get("INFO"))
		c.SetRequest(c.Request().WithContext(ctx))
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	}, middlewares.JWTMiddleware())

	// For Subscriptions
	// e.GET("/subscriptions", func(c echo.Context) error {
	// 	srv.ServeHTTP(c.Response(), c.Request())
	// 	return nil
	// })

	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
