package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
}
