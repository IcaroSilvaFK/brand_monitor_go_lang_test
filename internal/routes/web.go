package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeWebRoutes(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Icaro",
		})
	})

}
