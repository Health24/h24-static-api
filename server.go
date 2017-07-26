package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vshaveyko/h24-static-api/config"
	"github.com/vshaveyko/h24-static-api/controllers"
)

const (
	ACTION_INDEX = iota
	ACTION_SHOW
	ACTION_CREATE
	ACTION_UPDATE
	ACTION_DESTROY
)

var AppSecret = []byte(config.AppSecret)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))

	// Public routes
	// publicRoutes := e.Group("")
	// publicRoutes.POST("/login", controllers.Login)

	// Restricted routes
	apiRoutes := e.Group("/api")
	// apiRoutes.Use(middleware.JWT(AppSecret))

	// Resources
	resource(apiRoutes, "/static_translations", controllers.StaticTranslationsController{}, []int{ACTION_INDEX})

	e.Logger.Fatal(e.Start(":1323"))
}

func resource(e *echo.Group, url string, controller controllers.CRUDController, resourceActions []int) {
	if resourceActions == nil {
		resourceActions = []int{ACTION_INDEX, ACTION_SHOW, ACTION_CREATE, ACTION_UPDATE, ACTION_DESTROY}
	}
	for _, action := range resourceActions {
		resourceAction(e, url, controller, action)
	}
}

func resourceAction(e *echo.Group, url string, controller controllers.CRUDController, action int) {
	switch action {
	case ACTION_INDEX:
		e.GET(url, controllers.Index(controller))
	case ACTION_CREATE:
		e.POST(url, controllers.Create(controller))
	case ACTION_SHOW:
		e.GET(singularUrl(url), controllers.Show(controller))
	case ACTION_UPDATE:
		e.PUT(singularUrl(url), controllers.Update(controller))
		e.PATCH(singularUrl(url), controllers.Update(controller))
	case ACTION_DESTROY:
		e.DELETE(singularUrl(url), controllers.Destroy(controller))
	}
}

func singularUrl(url string) string {
	return fmt.Sprintf("%s/:id", url)
}
