package main

import (
	"covid-reactor/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := api.Api{}
	v1 := r.Group("/api/v1")
	{
		v1.GET("/info", api.GetInfo)
	}
	r.Run()
}
