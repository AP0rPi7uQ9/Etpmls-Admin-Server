package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteWeb(r *gin.Engine)  {
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
}
