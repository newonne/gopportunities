package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// inicializa o router utilizando as configura
	router := gin.Default()

	// inicializa as rotas
	InitializeRoutes(router)

	addr := ":8080"
	router.Run(addr) // listen and serve on 0.0.0.0:8080
}
