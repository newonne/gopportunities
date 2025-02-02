package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1") // Fixed variable assignment

	v1.GET("/opening", func(c *gin.Context) { // Fixed function declaration
		c.JSON(http.StatusOK, gin.H{ // Fixed JSON response
			"msg": "GET Opening",
		})
	})

	v1.POST("/opening", func(c *gin.Context) { // Fixed function declaration
		c.JSON(http.StatusOK, gin.H{ // Fixed JSON response
			"msg": "POST t Opening",
		})
	})

	v1.DELETE("/opening", func(c *gin.Context) { // Fixed function declaration
		c.JSON(http.StatusOK, gin.H{ // Fixed JSON response
			"msg": "DELETE Opening",
		})
	})

	v1.PUT("/opening", func(c *gin.Context) { // Fixed function declaration
		c.JSON(http.StatusOK, gin.H{ // Fixed JSON response
			"msg": "UPDATE Opening",
		})
	})

	v1.GET("/openings", func(c *gin.Context) { // Fixed function declaration
		c.JSON(http.StatusOK, gin.H{ // Fixed JSON response
			"msg": "GET Openings",
		})
	})
}
