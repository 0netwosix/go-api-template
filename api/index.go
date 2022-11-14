package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Index")
}
