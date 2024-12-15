package main

import (
	"net/http"
	"school/config"
	"school/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {

	db := config.InitDB()

	// INIT ECHO
	e := echo.New()
	apiV1 := e.Group("api/v1/")

	// ROUTING
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	userController := controllers.NewUserController(db)
	apiV1.POST("user/create", userController.Create)
	apiV1.PUT("user/update/:id", userController.Update)
	apiV1.DELETE("user/delete/:id", userController.Delete)
	apiV1.GET("user/get_all", userController.GetAll)
	apiV1.GET("user/detail/:id", userController.GetById)

	// SERVER HOST
	e.Logger.Fatal(e.Start(":9080"))
}
