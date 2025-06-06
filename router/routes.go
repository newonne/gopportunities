package router

import (
	"github.com/gin-gonic/gin"
	"github.com/newonne/gopportunities/handler"
)

func InitializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	v1 := router.Group("/api/v1") // Fixed variable assignment

	v1.GET("/opening", handler.ShowOpeningHandler)
	v1.POST("/opening", handler.CreateOpeningHandler)
	v1.DELETE("/opening", handler.DeleteOpeningHandler)
	v1.PUT("/opening", handler.UpdateOpeningHandler)
	v1.GET("/openings", handler.ListOpeningsHandler)

}
