package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiangyu123/create_view/apis"
)

func main() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/updateview", apis.UpdateView)
	}

	router.Run(":8000")
}
