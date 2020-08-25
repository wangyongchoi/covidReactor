package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
}

func (c *Api) GetInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello, Covid-Reactor")
}
