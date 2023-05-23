package routes

import (
	"github.com/HenrySaldanha/Go.FlashCards/controllers"
	docs "github.com/HenrySaldanha/Go.FlashCards/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequest() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	r.Use(cors.Default())
	r.GET("/cards", controllers.GetAllCards)
	r.GET("/cards/:id", controllers.GetCardById)
	r.DELETE("/cards/:id", controllers.DeleteCard)
	r.POST("/cards", controllers.InsertCard)
	r.PUT("/cards/:id", controllers.UpdateCard)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}
