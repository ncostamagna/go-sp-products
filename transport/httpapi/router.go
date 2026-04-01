package httpapi

import (
	"github.com/gin-gonic/gin"
)

func New(e Endpoints) *gin.Engine {
	app := gin.Default()

    app.POST("/products", e.Store)
	app.GET("/products", e.GetAll)
	app.GET("/products/:id", e.GetById)
	app.PATCH("/products/:id", e.Update)
	app.DELETE("/products/:id", e.Delete)

	return app
}