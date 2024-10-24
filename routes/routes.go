package routes

import (
	"shortLink/controllers"
	"github.com/gin-gonic/gin"
)


func HandleRequest(){
	app := gin.Default()


	app.GET("/admin/list",controllers.ListAll)

    app.GET("/:name",controllers.ListByName)
	app.POST("/",controllers.CreateLink)
	app.PATCH("/:id",controllers.EditLink)
	app.DELETE("/:id",controllers.DeleteLink)

	app.Run(":3000")

}