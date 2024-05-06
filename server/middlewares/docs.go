package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectToDocs(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	if path == "/api/docs" || path == "/api/docs/" {
		ctx.Redirect(http.StatusFound, "/api/docs/index.html")
		return
	}
	ctx.Next()
}
